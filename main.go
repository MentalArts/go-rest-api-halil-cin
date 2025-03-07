package main

import (
	"go-rest-api-halil-cin/internal/config"
	"go-rest-api-halil-cin/internal/db"
	"go-rest-api-halil-cin/internal/models"
	"go-rest-api-halil-cin/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	db, err := db.InitDB(cfg)
	if err != nil {
		panic("F.. Couldn't connect to Database")
	}

	db.AutoMigrate(&models.Book{}, &models.Author{}, &models.Review{})

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
