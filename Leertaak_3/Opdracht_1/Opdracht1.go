package main

import (
	"fmt"
)

func main() {
	// Data declaration part:
	a := 2
	b := 40
	var c float32 = 2
	var d float32 = 40
	// Print the result of the following operations:
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println("Operations complete.")
	// Print the result of the following operations:
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	fmt.Println(int(c) % int(d))
	fmt.Println("Operations complete.")
	// Efficiëntste notatie om 5 op te tellen:
	c = c + 5
	// Efficiëntste notatie om 1 op te tellen:
	c++
	// Call the function calculateOutput with 2 float32's:
	calculateOutput(1, 5)
	fmt.Println("Operations complete.")
	calculateOutput(3, 0)
	fmt.Println("Operations complete.")
	calculateOutput(3, 0.001)
	fmt.Println("Operations complete.")
	calculateOutput(7, -5)
	fmt.Println("Operations complete.")
}

// calculateOutput takes 2 float32 variables, performs mathematical operations on them and prints the result:
func calculateOutput(e, f float32) {
	fmt.Println(e*e + f*f)
	fmt.Println(e * (e + f) * f)
	fmt.Println((e + f) * (e - f))
	fmt.Println((e + f) * (e - f) / f)
}
