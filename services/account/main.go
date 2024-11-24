//go:generate protoc ./proto/account.proto --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	"github.com/dreadster3/pawcare/services/account/endpoint"
	"github.com/dreadster3/pawcare/services/account/proto"
	"github.com/dreadster3/pawcare/services/account/service"
	"github.com/dreadster3/pawcare/services/account/transport"
	"github.com/go-kit/log"
	kitlog "github.com/go-kit/log"
)

const (
	defaultHttpPort = "8080"
	defaultGrpcPort = "8081"
)

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

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
	var (
		httpPort = envString("HTTP_PORT", defaultHttpPort)
		grpcPort = envString("GRPC_PORT", defaultGrpcPort)

		httpAddr = flag.String("http.addr", ":"+httpPort, "HTTP listen address")
		grpcAddr = flag.String("grpc.addr", ":"+grpcPort, "gRPC listen address")
	)

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	svc := service.NewProfileService(logger)
	endpoints := endpoint.NewSet(svc, logger)

	httpHandler := transport.MakeHTTPHandler(endpoints, logger)
	grpcHandler := transport.NewGRPCServer(endpoints, logger)

	var g group.Group

	{
		logger := log.With(logger, "transport", "http")
		httpListenAddr, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			return err
		}

		g.Add(func() error {
			logger.Log("msg", "Starting server", "addr", *httpAddr)
			return http.Serve(httpListenAddr, httpHandler)
		}, func(err error) {
			logger.Log("msg", "Closing server", "reason", err)
			httpListenAddr.Close()
		})
	}

	{
		logger := log.With(logger, "transport", "grpc")
		grpcListenAddr, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			return err
		}

		g.Add(func() error {
			server := grpc.NewServer()
			proto.RegisterAccountServiceServer(server, grpcHandler)
			logger.Log("msg", "Starting server", "addr", *grpcAddr)
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
