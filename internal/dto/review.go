package dto

type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment" binding:"required"`
	BookID  uint   `json:"book_id" binding:"required"`
}

type UpdateReviewRequest struct {
	Rating  int    `json:"rating" binding:"min=1,max=5"`
	Comment string `json:"comment"`
}

type ReviewResponse struct {
	ID         uint   `json:"id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
	DatePosted string `json:"date_posted"`
	BookID     uint   `json:"book_id"`
}
