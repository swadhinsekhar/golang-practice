package main

import (
	"html/template"
	"log"
	"net/http"
)

//var template_html = template.Must(template.ParseGlob("template/*"))

func Home(w http.ResponseWriter, r *http.Request) {
	//var customers []Customer
	var template_html *template.Template

	//customers = GetCustomers()
	//log.Println(customers)

	template_html = template.Must(template.ParseFiles("main.html"))
	//template_html.ExecuteTemplate(w, nil)
	template_html.Execute(w, nil)
}

func main() {
	log.Println("Server started on : http://localhost:8000")

	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}
