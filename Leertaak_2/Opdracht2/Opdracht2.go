package main

import (
	"fmt"
)

func main() {
	// Scan the 2 variables:
	x := floatScanner("Vul uw eerste variabele in:")
	y := floatScanner("Vul uw tweede variabele in:")
	// Add the 2 variables and print the result:
	fmt.Println(x + y)
}

// floatScanner prints the query and returns the float64 scanned result:
func floatScanner(query string) (x float64) {
	fmt.Println(query)
	fmt.Scanln(&x)
	return x
}
