package routes

import (
	"go-rest-api-halil-cin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/books", handlers.GetBooks)
		api.GET("/books/:id", handlers.GetBook)
		api.POST("/books", handlers.CreateBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)

		api.GET("/authors", handlers.GetAuthors)
		api.GET("/authors/:id", handlers.GetAuthor)
		api.POST("/authors", handlers.CreateAuthor)
		api.PUT("/authors/:id", handlers.UpdateAuthor)
		api.DELETE("/authors/:id", handlers.DeleteAuthor)

		api.GET("/books/:id/reviews", handlers.GetReviewsForBook)
		api.POST("/books/:id/reviews", handlers.CreateReview)
		api.PUT("/reviews/:id", handlers.UpdateReview)
		api.DELETE("/reviews/:id", handlers.DeleteReview)
	}
}
