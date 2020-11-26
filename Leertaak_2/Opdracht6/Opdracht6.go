package main

import (
	"fmt"
)

func main() {
	// Print 5 sets of first and last name:
	for i := 0; i < 5; i++ {
		fmt.Println(printFullName(Scanner("Wat is uw voornaam?"), Scanner("Wat is uw achternaam?")))
	}
}

// printFullName returns the full name based on the first name and surname:
func printFullName(voornaam, achternaam string) (volledigeNaam string) {
	volledigeNaam = voornaam + " " + achternaam
	return volledigeNaam
}

// Scanner prints the query and returns the input string:
func Scanner(query string) (x string) {
	fmt.Println(query)
	fmt.Scan(&x)
	return x
}
