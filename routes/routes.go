package routes

import (
    "github.com/gorilla/mux"
    "my-knowledge-site/controllers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Routes pour les articles
    router.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
    router.HandleFunc("/articles", controllers.CreateArticle).Methods("POST")

    return router
}
