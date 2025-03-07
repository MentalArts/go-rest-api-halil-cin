package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name      string
	Biography string
	BirthDate string
	Books     []Book
}
