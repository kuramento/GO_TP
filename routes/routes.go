package routes

import (
	"my-knowledge-site/controllers"
	"github.com/gorilla/mux"
	"net/http" // Ajoutez cette ligne pour importer le package net/http
)

// Configuration des routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Définir les routes pour les articles
	router.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	router.HandleFunc("/articles", controllers.CreateArticle).Methods("POST")

	// Exemple d'une route par défaut (optionnel)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html") // Assurez-vous que index.html existe
	}).Methods("GET")

	// Retourner le routeur
	return router
}
