package main

import (
	"log"
	"net/http"
	"my-knowledge-site/routes"
)

func main() {
	// Configurer les routes
	router := routes.SetupRoutes()

	// Démarrer le serveur
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router) // Utilisez mux.Router pour servir les routes
	if err != nil {
		log.Fatal(err)
	}
}
