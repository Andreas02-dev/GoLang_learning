package main

import (
	"fmt"
)

func main() {
	// Create a stringarray and fill it with the 7 first elements and print it:
	elemNames := [7]string{"Waterstof", "Helium", "Lithium", "Beryllium", "Borium", "Koolstof", "Stikstof"}
	fmt.Println(elemNames)
	// Make a slice of integers using the make function:
	slice1 := make([]int, 11, 21)
	// Make a slice of strings using composite literals:
	slice2 := []string{}
	// Print slice1:
	fmt.Println(slice1)
	// Create and fill the following arrays with the correct values:
	countries := [10]string{"Netherlands", "United states", "United kingdom", "Australia", "India", "China", "Russia", "France", "Germany", "Spain"}
	top10Voornamen := [10]string{"Maria", "Johannes", "Johanna", "Jan", "Anna", "Cornelis", "Hendrik", "Cornelia", "Elisabeth", "Willem"}
	// Slice these arrays in 2 slices of equal length and print these:
	fmt.Println(countries[:5])
	fmt.Println(countries[5:])
	// Create a slice with only the woman names:
	slice2 = append(slice2, top10Voornamen[:1]...)
	slice2 = append(slice2, top10Voornamen[2:3]...)
	slice2 = append(slice2, top10Voornamen[4:5]...)
	slice2 = append(slice2, top10Voornamen[7:9]...)
	// Print the aforementioned slice:
	fmt.Println(slice2)
}
