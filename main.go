package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"microservices/handlers"
)

type handlerTuple struct {
	uri     string
	handler http.Handler
}

func main() {
	logger := createLogger()

	// Create handlers
	tuples := []handlerTuple{
		{"/products", handlers.NewProducts(logger)},
	}

	// Create new Serve Mux
	serveMux := http.NewServeMux()

	// Register handlers
	registerHandlers(tuples, serveMux)

	// Create a server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	logger.Println("received termination, graceful shutdown", sig)

	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second)
	server.Shutdown(timeoutCtx)
}

func createLogger() *log.Logger {
	return log.New(os.Stdout, "product-api", log.LstdFlags)
}

func registerHandlers(tuples []handlerTuple, serveMux *http.ServeMux) {
	for _, tuple := range tuples {
		serveMux.Handle(tuple.uri, tuple.handler)
	}
}
