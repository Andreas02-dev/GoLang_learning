package main

import (
	"fmt"
)

// Data declaration part:
var getal1 int8
var getal2 int8
var getal11 int
var getal22 int
var vermenigvuldig int

func main() {
	// Make both variables have the number of 100:
	getal1 = 100
	getal2 = 100
	// Add both numbers to variable totaal and print them to the terminal (gives overflow):
	totaal := getal1 + getal2
	// Convert both numbers to int and multiply them with each other:
	vermenigvuldig = int(getal1) * int(getal2)
	// Print totaal to the terminal:
	fmt.Println(totaal)
	// Print vermenigvuldig to the terminal:
	fmt.Println(vermenigvuldig)
}
