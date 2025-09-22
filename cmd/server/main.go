// Package main SCS Mission Service API
// @title SCS Mission Service API
// @version 1.0
// @description This is the SCS Mission Service API for managing security incidents and guidance procedures.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	config "scs-guard/config"
	"scs-guard/internal/container"
	"syscall"
	"time"

	"scs-guard/internal/server"
	"scs-guard/pkg/db"
	minio_client "scs-guard/pkg/minio"

	"scs-guard/pkg/logger"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func main() {
	// Load configuration from config file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	//Init logger
	appLogger := logger.GetLogger()
	appLogger.InitLogger(&cfg)
	appLogger.Infof("LogLevel: %s, Mode: %s", cfg.Logger.Level, cfg.Server.Mode)

	//Init db
	psqlDb, err := db.NewGormDB(&cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Info("Postgres connected")
	}

	if err != nil {
		appLogger.Fatalf("Database migration failed: %s", err)
	}
	//Init minio client
	minioClient, err := minio_client.NewMinioClient(cfg.Minio.Endpoint, cfg.Minio.AccessKey, cfg.Minio.SecretKey, cfg.Minio.BucketName, appLogger)
	if err != nil {
		appLogger.Fatalf("Minio init: %s", err)
	} else {
		appLogger.Info("Minio connected")
	}

	// Create shared repositories and services using container
	deps := container.NewContainer(psqlDb, minioClient)

	// Initialize the server with shared dependencies
	s := server.NewServer(&cfg, psqlDb, appLogger, deps)

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		if err := s.Run(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatalf("Error starting server: %v", err)
		}
	}()
	// Block until a signal is received
	<-quit

	appLogger.Info("Shutting down the server and consumer...")

	// Create a separate, timeout context for the server shutdown
	serverShutdownCtx, serverShutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer serverShutdownCancel()

	// Shut down the Echo server
	if err := s.Shutdown(serverShutdownCtx); err != nil {
		appLogger.Errorf("Server shutdown failed: %v", err)
	}

	appLogger.Info("Server and consumer stopped.")
}
