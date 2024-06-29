package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	// Define the /health endpoint
	router.HandleFunc("/health", HealthHandler).Methods("GET")

	// Define the / endpoint
	router.HandleFunc("/", RootHandler).Methods("GET")

	// Start the server
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}

// HealthHandler handles the /health endpoint
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "OK",
		Message: "Server is alive",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RootHandler handles the / endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "OK",
		Message: "Server is working",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
