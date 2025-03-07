package handlers

import (
	"go-rest-api-halil-cin/internal/dto"
	"go-rest-api-halil-cin/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func GetBooks(c *gin.Context) {
	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := db.Preload("Author").Preload("Reviews").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
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

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
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

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}
