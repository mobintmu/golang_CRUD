package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mobintmu/golang_crud/internal/repositories"
	"github.com/mobintmu/golang_crud/internal/routes"
	"github.com/mobintmu/golang_crud/internal/server"
	"github.com/mobintmu/golang_crud/internal/services"
	"go.uber.org/fx"
)

func main() {
	// Create a base context for the application
	ctx := context.Background()

	app := fx.New(
		fx.Provide(
			func() (*server.Config, error) {
				return server.LoadConfig("")
			},
			routes.SetupRouter,
			services.NewUserService,
			repositories.NewUserRepository,
			server.NewServer,
		),
		fx.Invoke(func(
			lc fx.Lifecycle,
			router *gin.Engine,
			server *server.Server,
		) error {
			return server.StartServer(lc, router)
		}),
	)

	// Start the application
	if err := app.Start(ctx); err != nil {
		log.Printf("Failed to start application: %v", err)
		fmt.Fprintf(os.Stderr, "Error starting application: %v\n", err)
		os.Exit(1)
	}

	// Log successful startup
	log.Println("Application is running...")

	// Set up signal handling
	go func() {
		// Create a channel for signals
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Wait for signal
		sig := <-sigChan
		log.Printf("Received signal %v, shutting down...", sig)

		// Create a context with timeout for graceful shutdown
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := app.Stop(shutdownCtx); err != nil {
			log.Printf("Error stopping application: %v", err)
			os.Exit(1)
		}

		log.Println("Application stopped successfully")
		os.Exit(0)
	}()

	// Block main goroutine to keep application running
	select {}
}
