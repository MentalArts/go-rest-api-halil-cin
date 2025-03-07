package dto

type CreateAuthorRequest struct {
	Name      string `json:"name" binding:"required"`
	Biography string `json:"biography" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}

type UpdateAuthorRequest struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}

type AuthorResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}
