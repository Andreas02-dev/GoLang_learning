package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Get the input from the scanner and put it in the variables:
	x := Scanner("Vul het eerste getal in:")
	y := Scanner("Vul het tweede getal in:")
	// Convert the values to xx and yy for int:
	xx, _ := strconv.Atoi(x)
	yy, _ := strconv.Atoi(y)
	// Print the adding of integers:
	fmt.Println(xx + yy)
	// Print the adding of strings:
	fmt.Println(x + y)
}

// Scanner function takes a query and returns the scanned input:
func Scanner(query string) (x string) {
	fmt.Println(query)
	fmt.Scan(&x)
	return x
}
