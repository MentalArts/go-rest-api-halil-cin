package handlers

import (
	"go-rest-api-halil-cin/internal/dto"
	"go-rest-api-halil-cin/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetReviewsForBook godoc
// @Summary Get reviews for a specific book
// @Description Get a list of reviews for a book by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {array} dto.ReviewResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/books/{id}/reviews [get]
func GetReviewsForBook(c *gin.Context) {
	bookID := c.Param("id")

	var reviews []models.Review
	if err := db.Where("book_id = ?", bookID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews for the book", "We honestly dont know why": err.Error()})
		return
	}
	if len(reviews) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No reviews found for this book"})
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

// CreateReview godoc
// @Summary Create a new review
// @Description Create a new review for a book
// @Tags reviews
// @Accept json
// @Produce json
// @Param review body dto.CreateReviewRequest true "Create review"
// @Success 201 {object} dto.ReviewResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/books/{id}/reviews [post]
func CreateReview(c *gin.Context) {
	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "Calm down champ.": err.Error()})
		return
	}
	if req.Rating < 1 || req.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5 Do you want me to rate you -20??"})
		return
	}

	review := models.Review{
		Rating:     req.Rating,
		Comment:    req.Comment,
		BookID:     req.BookID, // Use req.BookID directly
		DatePosted: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review", "Maybe you are not worthy?": err.Error()})
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

// UpdateReview godoc
// @Summary Update a review
// @Description Update a review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Param review body dto.UpdateReviewRequest true "Update review"
// @Success 200 {object} dto.ReviewResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/reviews/{id} [put]
func UpdateReview(c *gin.Context) {
	reviewID := c.Param("id")

	var req dto.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "Slow down": err.Error()})
		return
	}

	var review models.Review
	if err := db.First(&review, reviewID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found", "There is not...": err.Error()})
		return
	}

	if req.Rating != 0 {
		review.Rating = req.Rating
	}
	if req.Comment != "" {
		review.Comment = req.Comment
	}

	if err := db.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review", "We like it the way it is": err.Error()})
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

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review by its ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} map[string]string{"message": "Freedom of speech purged successfully"}
// @Failure 404 {object} map[string]string{"error": "Review could not be found"}
// @Failure 500 {object} map[string]string{"error": "Server got smacked"}
// @Router /api/v1/reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	reviewID := c.Param("id")

	var review models.Review
	if err := db.First(&review, reviewID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review could not be found"})
		return
	}

	if err := db.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review", "cant cancel everyone you know..": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Freedom of speech purged successfully"})
}
