package models

import "gorm.io/gorm"

// Modèle Article
type Article struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
