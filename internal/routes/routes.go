package routes

import (
	"go-rest-api-halil-cin/internal/handlers"
	"go-rest-api-halil-cin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// Public routes (accessible without authentication)
	api := router.Group("/api/v1")
	{
		// Books
		api.GET("/books", handlers.GetBooks)
		api.GET("/books/:id", handlers.GetBook)
		api.POST("/books", handlers.CreateBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", middleware.AdminOnly(), handlers.DeleteBook)

		// Authors
		api.GET("/authors", handlers.GetAuthors)
		api.GET("/authors/:id", handlers.GetAuthor)
		api.POST("/authors", handlers.CreateAuthor)
		api.PUT("/authors/:id", handlers.UpdateAuthor)
		api.DELETE("/authors/:id", middleware.AdminOnly(), handlers.DeleteAuthor)

		// Reviews
		api.GET("/books/:id/reviews", handlers.GetReviewsForBook)
		api.POST("/books/:id/reviews", handlers.CreateReview)
		api.PUT("/reviews/:id", handlers.UpdateReview)
		api.DELETE("/reviews/:id", middleware.AdminOnly(), handlers.DeleteReview)
	}

	// Auth routes (for registration, login, and token refresh)
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh-token", handlers.RefreshToken)
	}
}
