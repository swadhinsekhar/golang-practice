package main

import (
	"fmt"
	"log"
	"net/http"
)

func otherHandler(w http.ResponseWriter, r *http.Request) {

	for k, cookie := range r.Cookies() {
		fmt.Printf(" %v -> %s\n", k, cookie)
	}
	fmt.Fprintf(w, "otherHandler")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name: "token",
		Value:  "12345",
		MaxAge: 1000,
	}
	http.SetCookie(w, &c)
	fmt.Fprintf(w, "RootHandler")
}

func setupRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api", otherHandler)
}

func main() {
	fmt.Println("server started at 127.0.0.1:8777")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8777", nil))
}
