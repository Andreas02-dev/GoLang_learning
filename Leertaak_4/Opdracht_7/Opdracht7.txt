package main

import (
	"fmt"
)

// Data declaration part:
var dagenVanDeWeek []string
var neerslagPerDag []int
var windRichting []string = []string{"N", "N", "N", "N", "N", "N", "N"}
var multiDimensionale [][]int

func main() {
	// Fill the first slice in multiDimensionale with the minimum temperatures
	// And fill the second slice in multiDimensionale with the maximum temperatures:
	multiDimensionale = [][]int{{0, 0, -1, 2, 3, 1, -1}, {6, 4, 6, 6, 7, 7, 6}}
	// Fill the dagenVanDeWeek slice, neerslagPerDag slice and windRichting slice:
	dagenVanDeWeek = []string{"Maandag", "Dinsdag", "Woensdag", "Donderdag", "Vrijdag", "Zaterdag", "Zondag"}
	neerslagPerDag = []int{0, 5, 2, 3, 6, 8, 14}
	windRichting = []string{"N", "Z", "N", "O", "W", "W", "O"}
	// Find the minimum temperature of the week and print it:
	fmt.Println("Minimum:")
	fmt.Println(minimumValue(multiDimensionale[0]))
	// Find the maximum temperature of the week and print it:
	fmt.Println("Maximum:")
	fmt.Println(maximumValue(multiDimensionale[1]))
	// Find the average neerslag of the week and print it:
	fmt.Println("Gemiddelde neerslag:")
	fmt.Println(gemiddeldeNeerslag(neerslagPerDag))
	// Calculate the wind direction which has the most entries:
	mostEntries, tiedEntries := meestVoorkomende(windRichting)
	// Print the entry with the most entries:
	fmt.Printf("Most entries: %s\n", mostEntries)
	// Print the tied entries:
	fmt.Printf("Tied entries: %q\n", tiedEntries)
	// Append to all slices so they are 14 days long:
	neerslagPerDag = append(neerslagPerDag, 2, 5, 7, 2, 4, 9, 14)
	windRichting = append(windRichting, "N", "O", "Z", "N", "Z", "W", "W")
	multiDimensionale[0] = append(multiDimensionale[0], 0, 3, 6, 2, 4, 2, 4)
	multiDimensionale[1] = append(multiDimensionale[1], 7, 9, 10, 14, 12, 15, 17)
	// Delete the first 3 entries from all slices:
	neerslagPerDag = removeIntIndexFirst3(neerslagPerDag)
	windRichting = removeStringIndexFirst3(windRichting)
	multiDimensionale[0] = removeIntIndexFirst3(multiDimensionale[0])
	multiDimensionale[1] = removeIntIndexFirst3(multiDimensionale[1])
	// Find the minimum temperature of the week and print it:
	fmt.Println("Minimum:")
	fmt.Println(minimumValue(multiDimensionale[0]))
	// Find the maximum temperature of the week and print it:
	fmt.Println("Maximum:")
	fmt.Println(maximumValue(multiDimensionale[1]))
	// Find the average neerslag of the week and print it:
	fmt.Println("Gemiddelde neerslag:")
	fmt.Println(gemiddeldeNeerslag(neerslagPerDag))
	// Calculate the wind direction which has the most entries:
	mostEntries, tiedEntries = meestVoorkomende(windRichting)
	// Print the entry with the most entries:
	fmt.Printf("Most entries: %s\n", mostEntries)
	// Print the tied entries:
	fmt.Printf("Tied entries: %q\n", tiedEntries)
}

// minimumValue finds the minimum value in an int slice and returns it:
func minimumValue(slice []int) int {
	value := 100
	for _, num := range slice {
		if num <= value {
			value = num
		}
	}
	return value
}

// maximumValue finds the maximum value in an int slice and returns it:
func maximumValue(slice []int) int {
	value := 0
	for _, num := range slice {
		if num >= value {
			value = num
		}
	}
	return value
}

// gemiddeldeNeerslag takes a slice of ints and returns the average neerslag as an int:
func gemiddeldeNeerslag(slice []int) int {
	value := 0
	for _, num := range slice {
		value = value + num
	}
	return value / len(slice)
}

// meestVoorkomende takes a slice of strings and returns the entry which has the most entries.
// It also returns a string slice with the tied entries:
func meestVoorkomende(slice []string) (string, []string) {
	// Create a map to store the value and count:
	value := map[string]int{}
	// Count the amount of times a value exists:
	for _, v := range slice {
		value[v] = (value[v] + 1)
	}
	count := 0
	meestVoorkomendeString := ""
	tied := []string{}
	// See which key has the biggest value:
	for k, v := range value {
		if v > count {
			count = v
			meestVoorkomendeString = k
		}
	}
	// Check if any keys have the same value, if so put them in a stringslice:
	for k, v := range value {
		if count == v && k != meestVoorkomendeString {
			tied = append(tied, k)
		}
	}
	// Return the string with the entry which has the most values and a slice with the tied entries:
	return meestVoorkomendeString, tied
}

// removeStringIndex removes the first 3 entries in a string slice and returns the new string slice:
func removeStringIndexFirst3(slice []string) []string {
	returnSlice := append(slice[3:])
	return returnSlice
}

// removeIntIndex removes the first 3 entries in an int slice and returns the new int slice:
func removeIntIndexFirst3(slice []int) []int {
	returnSlice := append(slice[3:])
	return returnSlice
}
