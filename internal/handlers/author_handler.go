package handlers

import (
	"go-rest-api-halil-cin/internal/dto"
	"go-rest-api-halil-cin/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the input payload
// @Tags authors
// @Accept json
// @Produce json
// @Param author body dto.CreateAuthorRequest true "Create author"
// @Success 201 {object} dto.AuthorResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/authors [post]
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

// GetAuthors godoc
// @Summary Get all authors
// @Description Get a list of all authors
// @Tags authors
// @Accept json
// @Produce json
// @Success 200 {array} dto.AuthorResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/authors [get]
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

// GetAuthor godoc
// @Summary Get specific author info
// @Description Get an author by given id
// @Tags author
// @Accept json
// @Produce json
// @Success 200 {array} dto.AuthorResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/author/:id [get]
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
