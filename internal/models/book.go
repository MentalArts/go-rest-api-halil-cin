package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title           string
	AuthorID        uint
	Author          Author
	ISBN            string
	PublicationYear int
	Description     string
	Reviews         []Review
}
