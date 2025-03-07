package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating     int
	Comment    string
	DatePosted string
	BookID     uint
}
