package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Server represents an HTTP server
type Server struct {
	*http.Server
	config Config
}

// NewServer creates a new HTTP server instance
func NewServer(router *gin.Engine, config *Config) *Server {
	return &Server{
		Server: &http.Server{
			Addr:           config.Port,
			Handler:        router,
			ReadTimeout:    config.Timeout,
			WriteTimeout:   config.Timeout,
			MaxHeaderBytes: 1 << 20,
		},
		config: *config,
	}
}

// StartServer starts the HTTP server
func (s *Server) StartServer(lc fx.Lifecycle, router *gin.Engine) error {
	// Start server on application start
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Printf("Server failed to start: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctxShutdown, cancel := context.WithTimeout(ctx, s.config.Timeout)
			defer cancel()

			if err := s.Shutdown(ctxShutdown); err != nil {
				log.Printf("Server shutdown failed: %v", err)
				return err
			}
			log.Println("Server shut down successfully")
			return nil
		},
	})

	return nil
}
