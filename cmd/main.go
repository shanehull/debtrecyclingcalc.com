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
)

var (
	allowedOrigin = "*"
	serverHost    = "127.0.0.1"
	serverPort    = "8080"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	// Check if the SERVER_HOST env var is set and override
	envHost, ok := os.LookupEnv("SH_SERVER_HOST")
	if ok {
		serverHost = envHost
	}

	// Check if SH_ALLOWED_ORIGIN env var is set and override
	envOrigin, ok := os.LookupEnv("SH_ALLOWED_ORIGIN")
	if ok {
		allowedOrigin = envOrigin
		fmt.Printf("API routes allowed origin: %s\n", allowedOrigin)
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", cors(http.HandlerFunc(handlers.IndexHandler)))

	mux.HandleFunc("/calc", cors(http.HandlerFunc(handlers.CalcHandler)))

	// Run the server at
	serveAt := fmt.Sprintf("%s:%s", serverHost, serverPort)
	go func() {
		if err := http.ListenAndServe(serveAt, mux); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("API Server available at %s\n", serveAt)

	// Wait for interrupt signal.
	<-ctx.Done()

	// Sleep to ensure graceful shutdown
	fmt.Println("Server shutting down...")
	time.Sleep(5 * time.Second)

	// Return to default context.
	cancel()

	fmt.Println("Server stopped")
}

func cors(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, hx-target, hx-current-url, hx-request")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	}
}
