package main

import (
	"api-gateway/api"
	_ "api-gateway/cmd/docs"
	"api-gateway/internal/config"
	"api-gateway/service"
	"context"
	"crypto/tls"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title 		Person Finance Management API
// @version 	2.0
// @description This is a Person Finance Management server for Swagger integration.
// @schemes https
// @host 		localhost:4040
// @BasePath 	/api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authoriz
func main() {
	cfg := config.Load()

	addr := flag.String("addr", cfg.ApiGateway.Http.Host+":"+cfg.ApiGateway.Http.Port, "HTTP Server address")
	logger := slog.New(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
		}))
	flag.Parse()

	service, err := service.New(cfg)
	if err != nil {
		logger.Error("Failed to initialize services", "error", err)
		os.Exit(1)
	}

	router := api.Router(api.Options{Service: service})

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      *addr,
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	go func() {
		logger.Info("Starting server", "addr", srv.Addr)
		if err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}

	logger.Info("Server exiting")

}
