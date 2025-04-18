package main

import (
	"context"
	"fmt"
	"golang_crud/config"
	"golang_crud/internal/app/handlers"
	"golang_crud/internal/pkg/health/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	// Initialize health service
	healthService := service.NewHealthService()
	handler := handlers.NewHealthHandler(healthService)

	// Create HTTP server
	mux := http.NewServeMux()
	handler.SetupRoutes(mux)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a separate goroutine
	go func() {
		log.Printf("health check server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server:", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown server:", err)
	}

	log.Println("server shut down")
}
