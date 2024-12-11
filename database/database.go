package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"my-knowledge-site/models"
)

var DB *gorm.DB

// Connexion à la base de données SQLite
func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("knowledge.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Échec de la connexion à la base de données:", err)
	}

	// Migrer les modèles
	DB.AutoMigrate(&models.Article{})
}
