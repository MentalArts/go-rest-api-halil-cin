package routes

import (
	"go-rest-api-halil-cin/internal/handlers"
	"go-rest-api-halil-cin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh-token", handlers.RefreshToken)
	}

	// Public GET routes (accessible without authentication)
	public := router.Group("/api/v1")
	{
		// Books
		public.GET("/books", handlers.GetBooks)
		public.GET("/books/:id", handlers.GetBook)

		// Authors
		public.GET("/authors", handlers.GetAuthors)
		public.GET("/authors/:id", handlers.GetAuthor)

		// Reviews
		public.GET("/books/:id/reviews", handlers.GetReviewsForBook)
	}

	// Protected routes (require authentication)
	api := router.Group("/api/v1")
	api.Use(middleware.AuthRequired())
	{
		// Books
		api.POST("/books", middleware.AdminOnly(), handlers.CreateBook)
		api.PUT("/books/:id", middleware.AdminOnly(), handlers.UpdateBook)
		api.DELETE("/books/:id", middleware.AdminOnly(), handlers.DeleteBook)

		// Authors
		api.POST("/authors", middleware.AdminOnly(), handlers.CreateAuthor)
		api.PUT("/authors/:id", middleware.AdminOnly(), handlers.UpdateAuthor)
		api.DELETE("/authors/:id", middleware.AdminOnly(), handlers.DeleteAuthor)

		// Reviews
		api.POST("/books/:id/reviews", handlers.CreateReview)
		api.PUT("/reviews/:id", handlers.UpdateReview)
		api.DELETE("/reviews/:id", middleware.AdminOnly(), handlers.DeleteReview)
	}
}
