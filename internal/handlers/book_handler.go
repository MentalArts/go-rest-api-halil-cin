package handlers

import (
	"encoding/json"
	"go-rest-api-halil-cin/internal/cache"
	"go-rest-api-halil-cin/internal/dto"
	"go-rest-api-halil-cin/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Create book"
// @Success 201 {object} dto.BookResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/books [post]

func CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Details: err.Error(),
		})
		return
	}

	book := models.Book{
		Title:           req.Title,
		AuthorID:        req.AuthorID,
		ISBN:            req.ISBN,
		PublicationYear: req.PublicationYear,
		Description:     req.Description,
	}

	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		AuthorID:        book.AuthorID,
		ISBN:            book.ISBN,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
	})
}

// GetBooks godoc
// @Summary Get all books
// @Description Get a list of all books in the system
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} dto.BookResponse
// @Success 200 {object} map[string]string "There are no books in the system"
// @Failure 500 {object} map[string]string
// @Router /api/v1/books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve books",
			Details: err.Error(),
		})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "There are no books in the system"})
		return
	}

	var response []dto.BookResponse
	for _, book := range books {
		response = append(response, dto.BookResponse{
			ID:              book.ID,
			Title:           book.Title,
			AuthorID:        book.AuthorID,
			ISBN:            book.ISBN,
			PublicationYear: book.PublicationYear,
			Description:     book.Description,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetBook godoc
// @Summary Get a specific book
// @Description Get a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")

	// Check if the book is cached
	cachedBook, err := cache.Get("book:" + id)
	if err == nil {
		var book dto.BookResponse
		if err := json.Unmarshal([]byte(cachedBook), &book); err == nil {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	// Fetch the book from the database
	var book models.Book
	if err := db.Preload("Author").Preload("Reviews").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Book not found",
			Details: "The book with the given ID does not exist",
		})
		return
	}

	// Prepare the response
	response := dto.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		AuthorID:        book.AuthorID,
		ISBN:            book.ISBN,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
	}

	// Cache the book for 5 minutes
	if jsonData, err := json.Marshal(response); err == nil {
		cache.Set("book:"+id, jsonData, 5*time.Minute)
	}

	c.JSON(http.StatusOK, response)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dto.UpdateBookRequest true "Update book"
// @Success 200 {object} dto.BookResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Details: err.Error(),
		})
		return
	}

	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Book not found",
			Details: "The book with the given ID does not exist",
		})
		return
	}

	if req.Title != "" {
		book.Title = req.Title
	}
	if req.AuthorID != 0 {
		book.AuthorID = req.AuthorID
	}
	if req.ISBN != "" {
		book.ISBN = req.ISBN
	}
	if req.PublicationYear != 0 {
		book.PublicationYear = req.PublicationYear
	}
	if req.Description != "" {
		book.Description = req.Description
	}

	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update book",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		AuthorID:        book.AuthorID,
		ISBN:            book.ISBN,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
	})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} map[string]string "Ignorance is BLISS"
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Book not found",
			Details: "The book with the given ID does not exist",
		})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete book",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ignorance is bliss"})
}
