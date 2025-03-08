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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "Use proper input and try again": err.Error()})
		return
	}

	author := models.Author{
		Name:      req.Name,
		Biography: req.Biography,
		BirthDate: req.BirthDate,
	}

	if err := db.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author", "Check if its already exist": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors", "Check if they exist": err.Error()})
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
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} dto.AuthorResponse
// @Failure 404 {object} map[string]string{"error": "There is no Author with such credentials"}
// @Failure 500 {object} map[string]string{"error": "Internal server error"}
// @Router /api/v1/authors/{id} [get]
func GetAuthor(c *gin.Context) {
	id := c.Param("id")

	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no Author with such credentials"})
		return
	}

	c.JSON(http.StatusOK, dto.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	})
}

// UpdateAuthor godoc
// @Summary Update an author
// @Description Update an author with the input payload
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Param author body dto.UpdateAuthorRequest true "Update author"
// @Success 200 {object} dto.AuthorResponse
// @Failure 400 {object} map[string]string{"error": "Invalid Request"}
// @Failure 404 {object} map[string]string{"error": "There is no Author with such credentials"}
// @Failure 500 {object} map[string]string{"error": "Internal server error"}
// @Router /api/v1/authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "Try it with proper input": err.Error()})
		return
	}

	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no Author with such credentials"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author", "": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	})
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Description Delete an author by given id
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} map[string]string{"message": "Author deleted, Long live oppresive press"}
// @Failure 404 {object} map[string]string{"error": "There is no Author with such credentals"}
// @Failure 500 {object} map[string]string{"error": "Failed to delete Author"}
// @Router /api/v1/authors/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no author with such credentials"})
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author", "": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted, Long live oppresive press"})
}
