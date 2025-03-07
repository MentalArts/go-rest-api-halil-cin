package routes

import (
	"go-rest-api-halil-cin/internal/handlers"
	"go-rest-api-halil-cin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	//Rate limiter
	router.Use(middleware.RateLimiter())
	// Public routes
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh-token", handlers.RefreshToken)
	}

	// Protected routes
	api := router.Group("/api/v1")
	api.Use(middleware.AuthRequired())
	{
		// Book endpoints
		api.GET("/books", handlers.GetBooks)
		api.GET("/books/:id", handlers.GetBook)
		api.POST("/books", middleware.AdminOnly(), handlers.CreateBook)
		api.PUT("/books/:id", middleware.AdminOnly(), handlers.UpdateBook)
		api.DELETE("/books/:id", middleware.AdminOnly(), handlers.DeleteBook)

		// Author endpoints
		api.GET("/authors", handlers.GetAuthors)
		api.GET("/authors/:id", handlers.GetAuthor)
		api.POST("/authors", middleware.AdminOnly(), handlers.CreateAuthor)
		api.PUT("/authors/:id", middleware.AdminOnly(), handlers.UpdateAuthor)
		api.DELETE("/authors/:id", middleware.AdminOnly(), handlers.DeleteAuthor)
		// Reviews
		api.GET("/books/:id/reviews", handlers.GetReviewsForBook)
		api.POST("/books/:id/reviews", handlers.CreateReview)
		api.PUT("/reviews/:id", middleware.AdminOnly(), handlers.UpdateReview)
		api.DELETE("/reviews/:id", middleware.AdminOnly(), handlers.DeleteReview)
	}
}
