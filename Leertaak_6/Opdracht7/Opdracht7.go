package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// HandleFunc handles the function handleAanbevelingen:
	http.HandleFunc("/aanbevelingen", handleAanbevelingen)
	// HandleFUnc handles the function handleNieuweAanbeveling:
	http.HandleFunc("/nieuwe-aanbeveling", handleNieuweAanbeveling)
	// Listen and Serve the request:
	http.ListenAndServe("localhost:8080", nil)
}

// handleWelkom renders the html template:
func handleAanbevelingen(response http.ResponseWriter, request *http.Request) {
	render, _ := template.ParseFiles("aanbevelingen.html")
	// Render the page:
	render.Execute(response, nil)
}

// handleNieuwe parses the form and logs it:
func handleNieuweAanbeveling(response http.ResponseWriter, request *http.Request) {
	// Parse the form:
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	// Log the parsed values:
	log.Println(request.FormValue("name"))
	log.Println(request.FormValue("email"))
	log.Println(request.FormValue("comment"))
}
