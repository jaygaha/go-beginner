package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_4_go-kit/feedback"
)

/*
Go-Kit
-> is a bundle of different libraries that can use for building microservices
-> these libraries provide components for logging, metrics, tracing, rate limit, etc.

3 types of components:
- Endpoint: a function that takes a request and returns a response
- Transport: a way to send requests and receive responses
- Service: a function that takes a request and returns a response
*/

func main() {
	var httpAddr = flag.String("http", ":8800", "HTTP listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "feedback",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	ctx := context.Background()
	var srv feedback.Service
	{
		repository := feedback.NewRepo(logger)
		srv = feedback.NewService(repository, logger)
	}

	errs := make(chan error)
	// intercepts ctrl+c signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := feedback.MakeServerEndpoints(srv)
	// HTTP transport
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := feedback.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
