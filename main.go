package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-rest-api-halil-cin/internal/cache"
	"go-rest-api-halil-cin/internal/config"
	"go-rest-api-halil-cin/internal/db"
	"go-rest-api-halil-cin/internal/handlers"
	"go-rest-api-halil-cin/internal/middleware"
	"go-rest-api-halil-cin/internal/models"
	"go-rest-api-halil-cin/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	cfg := config.LoadConfig()

	// Initialize WHY NOT INITIALISE???
	cache.InitializeRedis("redis:6379", "", 0)
	log.Info("Redis initialized")

	// Initialize database
	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Info("Database connected")

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database")
	}
	log.Info("Database migrated")

	handlers.InitDB(db)

	// Set up Gin router
	router := gin.Default()
	router.Use(middleware.RateLimiter()) // Apply rate limiting
	routes.SetupRoutes(router)

	// Start server with graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Info("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Info("Server exited gracefully")
}
