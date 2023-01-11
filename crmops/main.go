package main

import (
	"log"
	"net/http"
	"text/template"
)

var template_html = template.Must(template.ParseGlob("templates/*"))

/*
func Home(w http.ResponseWriter, r *http.Request) {
	var customers = []Customer {
		Customer {
			CustomerId: 101,
			CustomerName: "Swadhin",
			SSN: "123456",
		},
		Customer {
			CustomerId: 100,
			CustomerName: "Romali",
			SSN: "2386343",
		},
	}
	var template_html *template.Template

	//customers = GetCustomers()
	log.Println(customers)

	template_html.ExecuteTemplate(w, "Home", customers)
}
*/

func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func main() {
	log.Println("Server started on : http://localhost:8000")

	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
	log.Println("Server Shutdown")
}
