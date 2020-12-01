package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// Set color to red and print the color name:
	color.Set(color.FgRed)
	fmt.Println("Red")
	// Set color to blue and print the color name:
	color.Set(color.FgBlue)
	fmt.Println("Blue")
	// Set color to magenta and print the color name:
	color.Set(color.FgMagenta)
	fmt.Println("Magenta")
	// Set color to yellow and print the color name:
	color.Set(color.FgYellow)
	fmt.Println("Yellow")
}
