package main

import (
	"fmt"
)

func main() {
	// Print the sum using the sum function:
	fmt.Println(sum(2, 3))
	// Print the full name using the printFullName function:
	printFullName("Andreas", "Hoornstra")
}

// sum takes 2 ints and returns the value of adding both together:
func sum(x, y int) (z int) {
	z = x + y
	return z
}

// printFullName prints the full name based on the first name and surname:
func printFullName(voornaam, achternaam string) {
	volledigeNaam := voornaam + " " + achternaam
	fmt.Println(volledigeNaam)
}
