package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start")
	router := mux.NewRouter()
	router.HandleFunc("/signin", abc).Methods("POST")

	fmt.Println("Listen and Server")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(corsObj)(router)))
}

func abc(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("great!"))
}

