package controllers

import (
	"my-knowledge-site/models"
	"my-knowledge-site/database"
	"net/http"
	"html/template"
	"encoding/json"
)

// Récupérer tous les articles et afficher un template HTML
func GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article
	database.DB.Find(&articles)

	// Charger le template HTML pour afficher les articles
	tmpl, err := template.ParseFiles("templates/articles.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données des articles
	tmpl.Execute(w, articles)
}

// Créer un nouvel article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Récupérer les données du formulaire
	var article models.Article
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sauvegarder l'article dans la base de données
	database.DB.Create(&article)

	// Répondre avec un message de succès
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}
