package main

import (
	"log"
	"net/http"

	"my-knowledge-site/routes"
	"my-knowledge-site/database"
)

func main() {
	// Initialiser la base de données
	database.Connect()

	// Charger les routes
	router := routes.SetupRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
