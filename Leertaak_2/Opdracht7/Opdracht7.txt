package main

import (
	"fmt"
)

func main() {
	// Print the sum using the sum function:
	fmt.Println(sum(2, 3))
}

// sum takes 2 ints and returns the value of adding both together:
func sum(x, y int) (z int) {
	z = x + y
	return z
}
