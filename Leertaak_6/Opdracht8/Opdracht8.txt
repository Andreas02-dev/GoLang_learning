package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Data declaration part:

// FormStruct is the form struct:
type FormStruct struct {
	Name    string
	Email   string
	Comment string
}

func main() {
	// HandleFunc handles the function handleNieuweAanbeveling:
	http.HandleFunc("/nieuwe-aanbeveling", handleNieuweAanbeveling)
	// HandleFunc handles the function handleAanbevelingen:
	http.HandleFunc("/aanbevelingen", handleAanbevelingen)
	// HandleFunc handles the function handleAanbeveling:
	http.HandleFunc("/aanbeveling", handleAanbeveling)
	// Listen to and Serve the request:
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}
}

// handleAanbevelingen renders the html template and writes the form data to "formdata.json":
func handleAanbevelingen(response http.ResponseWriter, request *http.Request) {
	// Parse the html template:
	render, _ := template.ParseFiles("aanbevelingen.html")
	// Parse the form:
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	// Create the FormHandler and fill it with the form values:
	FormHandler := FormStruct{
		Name:    request.FormValue("name"),
		Email:   request.FormValue("email"),
		Comment: request.FormValue("comment"),
	}
	// Marshal it to json:
	bs, err := json.Marshal(FormHandler)
	if err != nil {
		log.Println(err)
	}
	// Write it to the formdata.json file:
	ioutil.WriteFile("formdata.json", bs, 0655)
	// Render the page:
	render.Execute(response, nil)
}

// handleNieuweAanbeveling renders the form and sends it to "/aanbevelingen":
func handleNieuweAanbeveling(response http.ResponseWriter, request *http.Request) {
	// Parse the html template:
	render, _ := template.ParseFiles("nieuwe-aanbeveling.html")
	// Render the page:
	render.Execute(response, nil)
}

// handleAanbeveling reads "formdata.json" and renders the html template:
func handleAanbeveling(response http.ResponseWriter, request *http.Request) {
	// Parse the html template:
	render, _ := template.ParseFiles("aanbeveling.html")
	// Read "formdata.json":
	bs, err := ioutil.ReadFile("formdata.json")
	if err != nil {
		log.Println(err)
	}
	// Create FormHandler and unmarshal into it:
	FormHandler := FormStruct{}
	err = json.Unmarshal(bs, &FormHandler)
	if err != nil {
		log.Println(err)
	}
	// Render the page:
	err = render.Execute(response, FormHandler)
	if err != nil {
		log.Println(err)
	}
}
