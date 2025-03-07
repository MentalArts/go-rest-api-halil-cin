package main

import (
	"go-rest-api-halil-cin/internal/config"
	"go-rest-api-halil-cin/internal/db"
	"go-rest-api-halil-cin/internal/handlers"
	"go-rest-api-halil-cin/internal/models"
	"go-rest-api-halil-cin/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := db.InitDB(cfg)
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database")
	}

	handlers.InitDB(db)

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
