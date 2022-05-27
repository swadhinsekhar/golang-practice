package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler...")
	log.Println(r)
	fmt.Fprintf(w, "loginHandler")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("homePageHandler...")
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Couldn't render template", 500)
		return
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/login", loginHandler)
}

func main() {
	log.Println("Login API Server started...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
