package handlers

import (
	"go-rest-api-halil-cin/internal/dto"
	"go-rest-api-halil-cin/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReviewsForBook(c *gin.Context) {
	bookID := c.Param("id")

	var reviews []models.Review
	if err := db.Where("book_id = ?", bookID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.ReviewResponse
	for _, review := range reviews {
		response = append(response, dto.ReviewResponse{
			ID:         review.ID,
			Rating:     review.Rating,
			Comment:    review.Comment,
			DatePosted: review.DatePosted,
			BookID:     review.BookID,
		})
	}

	c.JSON(http.StatusOK, response)
}

func CreateReview(c *gin.Context) {
	bookID := c.Param("id")

	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Review{
		Rating:     req.Rating,
		Comment:    req.Comment,
		BookID:     req.BookID,
		DatePosted: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ReviewResponse{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
	})
}

func UpdateReview(c *gin.Context) {
	reviewID := c.Param("id")

	var req dto.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var review models.Review
	if err := db.First(&review, reviewID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review could not be found"})
		return
	}

	if req.Rating != 0 {
		review.Rating = req.Rating
	}
	if req.Comment != "" {
		review.Comment = req.Comment
	}

	if err := db.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.ReviewResponse{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
	})
}

func DeleteReview(c *gin.Context) {
	reviewID := c.Param("id")

	var review models.Review
	if err := db.First(&review, reviewID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review could not be found"})
		return
	}

	if err := db.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Freedom of speech purged successfuly"})
}
