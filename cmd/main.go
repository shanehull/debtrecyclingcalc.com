package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"debtrecyclingcalc.com/internal/buildinfo"
	"debtrecyclingcalc.com/internal/handlers"
	"debtrecyclingcalc.com/internal/middleware"
)

var (
	allowedOrigin   = "*"
	serverHost      = "127.0.0.1"
	serverPort      = "8080"
	htmxHash        = "sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+"
	hyperscriptHash = "sha384-PHB1Sh8oNP+x/7DnXnGgRL3pHqqjvJrrASslf5EClwgcJVQcmyf0fUqr6h39eO/t"
	echartsHash     = "sha384-pPi0zxBAoDu6+JXW/C68UZLvBUUtU+7zonhif43rqj7pxsGyqyqzcian2Rj37Rss"
	logger          *slog.Logger
)

func init() {
	logger = slog.New(
		slog.NewJSONHandler(
			os.Stdout, nil,
		),
	).With(slog.String("version", buildinfo.GitTag))
	slog.SetDefault(logger)
}

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	// If the SERVER_HOST env var is set, use that
	envHost, ok := os.LookupEnv("SERVER_HOST")
	if ok {
		serverHost = envHost
	}

	// If the ALLOWED_ORIGIN env var is set, use that
	envOrigin, ok := os.LookupEnv("ALLOWED_ORIGIN")
	if ok {
		allowedOrigin = envOrigin
		logger.Info("allowed origin set", "allowedOrigin", allowedOrigin)
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.Handle("/favicon.ico", fileServer)

	mux.HandleFunc("/",
		middleware.CSPMiddleware(
			middleware.CORS(
				http.HandlerFunc(
					handlers.IndexHandler,
				),
				allowedOrigin),
			htmxHash,
			hyperscriptHash,
			echartsHash,
		),
	)

	mux.HandleFunc("/calc",
		middleware.CSPMiddleware(
			middleware.CORS(
				http.HandlerFunc(
					handlers.CalcHandler,
				),
				allowedOrigin),
			htmxHash,
			hyperscriptHash,
			echartsHash,
		),
	)

	mux.HandleFunc("/healthz", http.HandlerFunc(handlers.HealthzHandler))

	// Run the server at
	serveAt := fmt.Sprintf("%s:%s", serverHost, serverPort)
	go func() {
		if err := http.ListenAndServe(serveAt, mux); err != nil {
			log.Fatal(err)
		}
	}()
	logger.Info("server listening", "serverHost", serverHost, "serverPort", serverPort)

	// Wait for interrupt signal.
	<-ctx.Done()

	// Sleep to ensure graceful shutdown
	sleepSeconds := 5
	logger.Info("shutting down", "sleepSeconds", sleepSeconds)
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Return to default context.
	cancel()

	logger.Info("server shut down gracefully")
}
