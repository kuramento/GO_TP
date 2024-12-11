package main

import (
	"encoding/json"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handleTasks)
	mux.HandleFunc("/tasks/", handleTaskByID)
	return mux
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllTasks(w, r)
	case http.MethodPost:
		CreateTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extraction et gestion de l'ID dans l'URL
}
