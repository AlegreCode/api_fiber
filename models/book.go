package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    uint
	Author      Author `json:"author"`
}
