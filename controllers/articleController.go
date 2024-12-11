package controllers

import (
	"encoding/json"
	"net/http"
	"my-knowledge-site/database"
	"my-knowledge-site/models"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article
	database.DB.Find(&articles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Create(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
