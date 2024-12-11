package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialisation de la base de données
	InitDatabase()

	// Démarrage du serveur
	log.Println("Starting TaskManager server on port 8080...")
	if err := http.ListenAndServe(":8080", Router()); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
