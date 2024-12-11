package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"my-knowledge-site/models"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("knowledge.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrer les modèles
	DB.AutoMigrate(&models.Article{})
}
