package main

import (
	"fmt"
)

// Data declaration part:
var dagenVanDeWeek [7]string
var neerslagPerDag [7]int
var windRichting [7]string = [7]string{"N", "N", "N", "N", "N", "N", "N"}
var multiDimensionale [2][7]int

func main() {
	// Print the type and the first & last entry of the array:
	fmt.Printf("%#v\n%#v\n", dagenVanDeWeek[0], dagenVanDeWeek[6])
	// Print the length of the dagenVanDeWeek array:
	fmt.Println(len(dagenVanDeWeek))
	// Print the neerslagPerDag array and its type:
	fmt.Printf("%#v\n", neerslagPerDag)
	// Print the windRichting array and its type:
	fmt.Printf("%#v\n", windRichting)
	// Print the multiDimensionale array and its type:
	fmt.Printf("%#v\n", multiDimensionale)
	// Fill the first array in multiDimensionale with the minimum temperatures
	// And fill the second array in multiDimensionale with the maximum temperatures:
	multiDimensionale[0] = [7]int{0, 0, -1, 2, 3, 1, -1}
	multiDimensionale[1] = [7]int{6, 4, 6, 6, 7, 7, 6}
	// Fill the dagenVanDeWeek array, neerslagPerDag array and windRichting array:
	dagenVanDeWeek = [7]string{"Maandag", "Dinsdag", "Woensdag", "Donderdag", "Vrijdag", "Zaterdag", "Zondag"}
	neerslagPerDag = [7]int{0, 5, 2, 3, 6, 8, 14}
	windRichting = [7]string{"N", "Z", "N", "O", "W", "W", "O"}
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
}

// minimumValue finds the minimum value in an int array of 7 and returns it:
func minimumValue(array [7]int) int {
	value := 100
	for _, num := range array {
		if num <= value {
			value = num
		}
	}
	return value
}

// maximumValue finds the maximum value in an int array of 7 and returns it:
func maximumValue(array [7]int) int {
	value := 0
	for _, num := range array {
		if num >= value {
			value = num
		}
	}
	return value
}

// gemiddeldeNeerslag takes an array of int7 and returns the average neerslag as an int:
func gemiddeldeNeerslag(array [7]int) int {
	value := 0
	for _, num := range array {
		value = value + num
	}
	return value / len(array)
}

// meestVoorkomende takes an array of string7 and returns the entry which has the most entries.
// It also returns a string array with the tied entries:
func meestVoorkomende(array [7]string) (string, []string) {
	// Create a map to store the value and count:
	value := map[string]int{}
	// Count the amount of times a value exists:
	for _, v := range array {
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
	// Return the string with the entry which has the most values and an array with the tied entries:
	return meestVoorkomendeString, tied
}
