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
	http.ListenAndServe("localhost:8080", nil)
}

// handleWelkom renders the html template:
func handleWelkom(response http.ResponseWriter, request *http.Request) {
	render, _ := template.ParseFiles("welkom.html")

	// Data declaration part:
	data := struct {
		Name       string
		Role       string
		Techniques []string
	}{
		Name: "Andreas Hoornstra",
		Role: "Software Developer",
		Techniques: []string{
			"Java",
			"Golang",
			"Python",
			"HTML",
			"CSS",
			"SAP ABAP",
		},
	}

	// Render the page with the data struct values:
	render.Execute(response, data)
}
