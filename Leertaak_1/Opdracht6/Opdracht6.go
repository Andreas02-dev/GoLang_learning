package main

import (
	"fmt"
)

// Data declaration part:

var ontbijtSlice []string = []string{"1", "Sinaasappelsap (fles)", "2", "Brood", "3", "Melk (1L)"}
var dinerSlice []string = []string{"1", "kipfilet", "1", "Spinazie (400g)", "2", "Vanilleyoghurt (1L)"}
var boodschappenlijstSlice []string = []string{"Aantal: ", "Wat: ", "Maaltijd: "}

func main() {
	// Calling the slice printing function on all the slices:
	slicePrint(boodschappenlijstSlice)
	slicePrint(ontbijtSlice)
	slicePrint(dinerSlice)
}

// Declaring the slice printing function:
func slicePrint(inputSlice []string) {
	// Iterate over inputSlice and print the output and a tab after:
	for i, output := range inputSlice {
		fmt.Printf(output + "\t")
		// If i is the same as the length -1, place a newline after:
		if i == (len(inputSlice) - 1) {
			fmt.Printf("\n")
		}
	}
}
