package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"debtrecyclingcalculator.com.au/internal/handlers"
	"debtrecyclingcalculator.com.au/internal/middleware"
)

var (
	allowedOrigin = "*"
	serverHost    = "127.0.0.1"
	serverPort    = "8080"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	// Check if the SERVER_HOST env var is set and override
	envHost, ok := os.LookupEnv("SERVER_HOST")
	if ok {
		serverHost = envHost
	}

	// Check if SH_ALLOWED_ORIGIN env var is set and override
	envOrigin, ok := os.LookupEnv("ALLOWED_ORIGIN")
	if ok {
		allowedOrigin = envOrigin
		fmt.Printf("Allowed origin: %s\n", allowedOrigin)
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/",
		middleware.CORS(
			http.HandlerFunc(
				handlers.IndexHandler,
			),
			allowedOrigin),
	)

	mux.HandleFunc("/calc",
		middleware.CORS(
			http.HandlerFunc(
				handlers.CalcHandler,
			),
			allowedOrigin),
	)

	mux.HandleFunc("/healthz", http.HandlerFunc(handlers.HealthzHandler))

	// Run the server at
	serveAt := fmt.Sprintf("%s:%s", serverHost, serverPort)
	go func() {
		if err := http.ListenAndServe(serveAt, mux); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Server available at %s\n", serveAt)

	// Wait for interrupt signal.
	<-ctx.Done()

	// Sleep to ensure graceful shutdown
	fmt.Println("Server shutting down...")
	time.Sleep(5 * time.Second)

	// Return to default context.
	cancel()

	fmt.Println("Server stopped")
}
