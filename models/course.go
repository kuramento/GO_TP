package models

import "gorm.io/gorm"

// Modèle Course
type Course struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
