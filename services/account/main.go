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

	"github.com/dreadster3/pawcare/services/account/internal/config"
	ownerendpoint "github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	ownerservice "github.com/dreadster3/pawcare/services/account/internal/owner/service"
	petendpoint "github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/services/account/internal/repository/mongo"
	"github.com/dreadster3/pawcare/services/account/internal/transport"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/db/mongodb"
	"github.com/go-kit/log"

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

	db, teardown, err := mongodb.ConnectDB(ctx, viper.GetString(common.DBConnectionStringKey), "accounts")
	defer teardown(ctx)
	if err != nil {
		return err
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	ownerRepository := mongo.NewOwnerRepository(db, log.With(logger, "repository", "owner"))
	petRepository := mongo.NewPetRepository(db)

	ownerService := ownerservice.NewOwnerService(ownerRepository, log.With(logger, "service", "owner"))
	petService := petservice.NewPetService(petRepository, ownerService, log.With(logger, "service", "pet"))

	ownerEndpoints := ownerendpoint.NewSet(viper, ownerService, log.With(logger, "endpoint", "owner"))
	petEndpoints := petendpoint.NewSet(viper, ownerService, petService, log.With(logger, "endpoint", "pet"))

	httpHandler := transport.MakeHTTPServer(ownerEndpoints, petEndpoints, log.With(logger, "transport", "http"))
	httpHandler = accessControl(httpHandler)

	var g group.Group

	{
		httpAddr := fmt.Sprintf(":%s", viper.GetString(common.HTTPPortKey))
		logger := log.With(logger, "transport", "http")
		httpListenAddr, err := net.Listen("tcp", httpAddr)
		if err != nil {
			return err
		}

		g.Add(func() error {
			logger.Log("msg", "Starting server", "addr", httpAddr)
			return http.Serve(httpListenAddr, httpHandler)
		}, func(err error) {
			logger.Log("msg", "Closing server", "reason", err)
			httpListenAddr.Close()
		})
	}

	{
		grpcAddr := fmt.Sprintf(":%s", viper.GetString(common.GRPCPortKey))
		logger := log.With(logger, "transport", "grpc")
		grpcListenAddr, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			return err
		}

		g.Add(func() error {
			server := transport.NewGRPCServer(ownerEndpoints, petEndpoints, logger)
			logger.Log("msg", "Starting server", "addr", grpcAddr)
			return server.Serve(grpcListenAddr)
		}, func(err error) {
			logger.Log("msg", "Closing server", "reason", err)
			grpcListenAddr.Close()
		})
	}

	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		return fmt.Errorf("%s", <-c)
	}, func(err error) {
		logger.Log("msg", "Shutdown signal received", "signal", err)
	})

	return g.Run()
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
