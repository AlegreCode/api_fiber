package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	Name     string `json:"name"`
	Lastname string `json:"lastname"`

	Books []Book `gorm:"foreignKey:AuthorID"`
}
