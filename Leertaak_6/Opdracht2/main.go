package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/welkom", handleWelkom)
	http.ListenAndServe(":8080", nil)
}

func handleWelkom(response http.ResponseWriter, request *http.Request) {
	render, _ := template.ParseFiles("welkom.html")
	render.Execute(response, nil)
}
