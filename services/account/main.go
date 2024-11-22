package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dreadster3/pawcare/services/account/endpoint"
	"github.com/dreadster3/pawcare/services/account/service"
	"github.com/dreadster3/pawcare/services/account/transport"
	kitlog "github.com/go-kit/log"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
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

func main() {
	var (
		addr = envString("PORT", defaultPort)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	svc := service.NewProfileService(logger)
	endpoints := endpoint.NewSet(svc, logger)
	handler := transport.MakeHTTPHandler(endpoints, logger)

	http.Handle("/", handler)

	errChan := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		logger.Log("msg", "Press Ctrl+C to stop")
		if err := http.ListenAndServe(*httpAddr, nil); err != nil {
			errChan <- err
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errChan)
}
