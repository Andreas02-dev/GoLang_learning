package main

import (
	"fmt"
	"strconv"
)

// Data declaration part:
var string1 string = "1234567"
var string2 string = "6543210"

func main() {
	// Convert the strings to ints:
	string11, _ := strconv.Atoi(string1)
	string22, _ := strconv.Atoi(string2)
	// Pass the printed string to s:
	s := fmt.Sprintf("%d is string 1, %d is string 2", string11, string22)
	// Pass the printed string to stringAdded:
	stringAdded := fmt.Sprintf("%d is string1+2", string11+string22)
	// Print s:
	fmt.Println(s)
	// Print the type of stringAdded:
	fmt.Printf("The type of stringAdded is %T\n", stringAdded)
	// Convert string11 and string22 to float64-holding variables fl1 and fl2:
	fl1 := float64(string11)
	fl2 := float64(string22)
	// Print the values of these variables:
	fmt.Printf("Float1: %f, Float2: %f\n", fl1, fl2)
	// Add certain values to fl1 and fl2 and pass it to flAdded-variables:
	flAdded1 := fl1 + 0.6543210
	flAdded2 := fl2 + 0.1234567
	// Print flAdded1 and flAdded2:
	fmt.Printf("After adding float1: %f, after adding float2: %f\n", flAdded1, flAdded2)
	// Convert the floats to int64s:
	in1 := int64(flAdded1)
	in2 := int64(flAdded2)
	// Print the ints:
	fmt.Println(in1, in2)
	// Convert the floats to strings:
	string111 := strconv.FormatFloat(flAdded1, 'f', 6, 64)
	string222 := strconv.FormatFloat(flAdded2, 'f', 6, 64)
	// Print the strings:
	fmt.Println(string111, string222)
	// Convert the strings back to floats:
	flo1, _ := strconv.ParseFloat(string111, 64)
	flo2, _ := strconv.ParseFloat(string222, 64)
	// Print the strings:
	fmt.Printf("Float1: %f, Float2: %f", flo1, flo2)
}
