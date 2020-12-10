package main

import (
	"html/template"
	"net/http"
)

func main() {
	// HandleFunc handles the function handleAanbevelingen:
	http.HandleFunc("/aanbevelingen", handleAanbevelingen)
	// Listen and Serve the request:
	http.ListenAndServe("localhost:8080", nil)
}

// handleWelkom renders the html template:
func handleAanbevelingen(response http.ResponseWriter, request *http.Request) {
	render, _ := template.ParseFiles("aanbevelingen.html")
	// Render the page:
	render.Execute(response, nil)
}
