package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"go.uber.org/zap"

	"github.com/dreadster3/pawcare/services/account/internal/config"
	ownerendpoint "github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	ownerservice "github.com/dreadster3/pawcare/services/account/internal/owner/service"
	petendpoint "github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/services/account/internal/repository/mongo"
	"github.com/dreadster3/pawcare/services/account/internal/transport"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/db/mongodb"
	"github.com/dreadster3/pawcare/shared/events"

	"github.com/joho/godotenv"
)

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func _main() error {
	godotenv.Load()

	viper := config.InitConfig()
	ctx := context.Background()

	http.DefaultTransport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = viper.GetBool(common.InsecureSkipVerifyKey)

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()

	db, teardown, err := mongodb.ConnectDB(ctx, viper.GetString(common.DBConnectionStringKey), "accounts")
	defer teardown(ctx)
	if err != nil {
		return err
	}

	kafkaDispatcher := events.NewKafkaEventDispatcher(viper.GetStringSlice(common.KafkaBrokersKey), "accounts")
	defer kafkaDispatcher.Close()

	ownerRepository := mongo.NewOwnerRepository(db, logger.With(zap.String("repository", "owner")))
	petRepository := mongo.NewPetRepository(db, logger.With(zap.String("repository", "pet")))

	ownerService := ownerservice.NewOwnerService(ownerRepository, logger.With(zap.String("service", "owner")))
	petService := petservice.NewPetService(petRepository, ownerService, kafkaDispatcher, logger.With(zap.String("service", "pet")))

	ownerEndpoints, err := ownerendpoint.NewSet(viper, ownerService)
	if err != nil {
		return err
	}

	petEndpoints, err := petendpoint.NewSet(viper, petService)
	if err != nil {
		return err
	}

	httpHandler := transport.MakeHTTPServer(ownerEndpoints, petEndpoints, logger.With(zap.String("transport", "http")))
	httpHandler = accessControl(httpHandler)

	var g group.Group

	{
		httpAddr := fmt.Sprintf(":%s", viper.GetString(common.HTTPPortKey))
		logger := logger.With(zap.String("transport", "http"))
		httpListenAddr, err := net.Listen("tcp", httpAddr)
		if err != nil {
			return err
		}

		g.Add(func() error {
			logger.Info("Starting server", zap.String("addr", httpAddr))
			return http.Serve(httpListenAddr, httpHandler)
		}, func(err error) {
			logger.Info("Closing server", zap.NamedError("reason", err))
			httpListenAddr.Close()
		})
	}

	{
		grpcAddr := fmt.Sprintf(":%s", viper.GetString(common.GRPCPortKey))
		logger := logger.With(zap.String("transport", "grpc"))
		grpcListenAddr, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			return err
		}

		g.Add(func() error {
			server := transport.NewGRPCServer(ownerEndpoints, petEndpoints, logger)
			logger.Info("Starting server", zap.String("addr", grpcAddr))
			return server.Serve(grpcListenAddr)
		}, func(err error) {
			logger.Info("Closing server", zap.NamedError("reason", err))
			grpcListenAddr.Close()
		})
	}

	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		return fmt.Errorf("%s", <-c)
	}, func(err error) {
		logger.Info("Shutdown signal received", zap.NamedError("signal", err))
	})

	return g.Run()
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
