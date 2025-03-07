package dto

type CreateBookRequest struct {
	Title           string `json:"title" binding:"required"`
	AuthorID        uint   `json:"author_id" binding:"required"`
	ISBN            string `json:"isbn" binding:"required"`
	PublicationYear int    `json:"publication_year" binding:"required"`
	Description     string `json:"description" binding:"required"`
}

type UpdateBookRequest struct {
	Title           string `json:"title"`
	AuthorID        uint   `json:"author_id"`
	ISBN            string `json:"isbn"`
	PublicationYear int    `json:"publication_year"`
	Description     string `json:"description"`
}

type BookResponse struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	AuthorID        uint   `json:"author_id"`
	ISBN            string `json:"isbn"`
	PublicationYear int    `json:"publication_year"`
	Description     string `json:"description"`
}
