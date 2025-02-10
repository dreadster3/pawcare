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

	"github.com/dreadster3/pawcare/services/account/config"
	ownerendpoint "github.com/dreadster3/pawcare/services/account/owner/endpoint"
	ownerservice "github.com/dreadster3/pawcare/services/account/owner/service"
	petendpoint "github.com/dreadster3/pawcare/services/account/pet/endpoint"
	petservice "github.com/dreadster3/pawcare/services/account/pet/service"
	"github.com/dreadster3/pawcare/services/account/repository/mongo"
	"github.com/dreadster3/pawcare/services/account/transport"
	"github.com/dreadster3/pawcare/shared/db/mongodb"
	kitlog "github.com/go-kit/log"

	"github.com/joho/godotenv"
)

const (
	DefaultHttpPort string = "8080"
	DefaultGrpcPort string = "8081"
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
	connectionString := BuildConnectionString(viper)
	ctx := context.Background()

	db, teardown, err := mongodb.ConnectDB(ctx, connectionString, "accounts")
	defer teardown(ctx)
	if err != nil {
		return err
	}

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	ownerRepository := mongo.NewOwnerRepository(db, kitlog.With(logger, "repository", "owner"))
	petRepository := mongo.NewPetRepository(db)

	ownerService := ownerservice.NewOwnerService(ownerRepository, kitlog.With(logger, "service", "owner"))
	petService := petservice.NewPetService(petRepository, ownerService, kitlog.With(logger, "service", "pet"))

	ownerEndpoints := ownerendpoint.NewSet(viper, ownerService, kitlog.With(logger, "endpoint", "owner"))
	petEndpoints := petendpoint.NewSet(viper, ownerService, petService, kitlog.With(logger, "endpoint", "pet"))

	httpHandler := transport.MakeHTTPServer(ownerEndpoints, petEndpoints, kitlog.With(logger, "transport", "http"))
	httpHandler = accessControl(httpHandler)

	var g group.Group

	{
		httpAddr := fmt.Sprintf(":%s", viper.GetString(config.HTTPPortKey))
		logger := kitlog.With(logger, "transport", "http")
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
		grpcAddr := fmt.Sprintf(":%s", viper.GetString(config.GRPCPortKey))
		logger := kitlog.With(logger, "transport", "grpc")
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
