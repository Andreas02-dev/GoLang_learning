package main

import (
	"fmt"
)

func main() {
	// Create an anonymous struct using a composite literal:
	v1 := struct {
		color string
	}{
		color: "Blue",
	}
	// Print the struct:
	fmt.Println(v1)
}
