package handlers

import (
	"go-rest-api-halil-cin/internal/dto"
	"go-rest-api-halil-cin/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func CreateAuthor(c *gin.Context) {
	var req dto.CreateAuthorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{
		Name:      req.Name,
		Biography: req.Biography,
		BirthDate: req.BirthDate,
	}

	if err := db.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	})
}

func GetAuthors(c *gin.Context) {
	var authors []models.Author

	if err := db.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.AuthorResponse
	for _, author := range authors {
		response = append(response, dto.AuthorResponse{
			ID:        author.ID,
			Name:      author.Name,
			Biography: author.Biography,
			BirthDate: author.BirthDate,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetAuthor(c *gin.Context) {
	id := c.Param("id")

	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	c.JSON(http.StatusOK, dto.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	})
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	if req.Name != "" {
		author.Name = req.Name
	}
	if req.Biography != "" {
		author.Biography = req.Biography
	}
	if req.BirthDate != "" {
		author.BirthDate = req.BirthDate
	}

	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, dto.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	})
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "author deleted"})
}
