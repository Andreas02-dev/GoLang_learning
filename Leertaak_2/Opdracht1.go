package main

import (
	"fmt"
)

func main() {
	// Storing the first and last name in the variables using the Scanner function:
	voornaam := Scanner("Wat is uw voornaam?")
	achternaam := Scanner("Wat is uw achternaam?")
	// Print the first and last name separated by space:
	fmt.Print(voornaam, " ", achternaam)
}

// Scanner function takes a query and returns the scanned input:
func Scanner(query string) (x string) {
	fmt.Println(query)
	fmt.Scan(&x)
	return x
}
