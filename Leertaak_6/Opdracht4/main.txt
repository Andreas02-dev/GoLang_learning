package main

import (
	"html/template"
	"net/http"
)

func main() {
	// HandleFunc handles the function handleWelkom:
	http.HandleFunc("/welkom", handleWelkom)
	// Make sure the page can handle images:
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	// Listen and Serve the request:
	http.ListenAndServe(":8080", nil)
}

// handleWelkom renders the html template:
func handleWelkom(response http.ResponseWriter, request *http.Request) {
	render, _ := template.ParseFiles("welkom.html")
	render.Execute(response, nil)
}
