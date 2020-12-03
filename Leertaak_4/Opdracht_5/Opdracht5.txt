package main

import (
	"fmt"
)

// Data declaration part:
var array [6]string
var slice []string

func main() {
	// FIll the array and slice with 5 names:
	array[0] = "Anniek"
	array[1] = "Andreas"
	array[2] = "Thomas"
	array[3] = "Peter"
	array[4] = "Margriet"
	slice = []string{"Elizabeth", "Ruud", "Paula", "Laurens", "Hannie"}
	// Print the array and slice with the type:
	fmt.Printf("%#v\n", array)
	fmt.Printf("%#v\n", slice)
	// Change the second name in both the array and the slice:
	array[1] = "Anniek2"
	slice[1] = "Anniek"
	// Print the array and slice with the type:
	fmt.Printf("%#v\n", array)
	fmt.Printf("%#v\n", slice)
	// An entry can't be removed from an array, however we can set it to "":
	array[2] = ""
	// Remove the third name from the slice using removeIndex:
	slice = removeIndex(slice, 2)
	// Print the array and slice with the type:
	fmt.Printf("%#v\n", array)
	fmt.Printf("%#v\n", slice)
	// Add an entry to the array and slice:
	array[5] = "Anniek again!"
	slice = append(slice, "Here I am!")
	// Print the array and slice with the type:
	fmt.Printf("%#v\n", array)
	fmt.Printf("%#v\n", slice)
}

// removeIndex takes a string slice and an int index and returns the slice without the entry at index:
func removeIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
