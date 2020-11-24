package main

import (
	"bufio"
	"fmt"
	"os"
)

// Data declaration part:
var maaltijdSlice []string = []string{"Sinaasappelsap (fles)", "1", "Ontbijt", "Brood", "2", "Ontbijt", "Melk (1L)", "3", "Ontbijt", "Kipfilet", "1", "Diner", "Spinazie (400g)", "1", "diner", "Vanilleyoghurt (1L)", "2", "Diner"}
var headerSlice []string = []string{"Wat: ", "Aantal: ", "Maaltijd: "}

func main() {
	// Ask what someone would like to buy and append it to the list:
	userInput()
	userInput()

	// Check which entry is longest and place spaces:
	checkLongestOne(maaltijdSlice, headerSlice)
	// Calling the slice printing function on all the slices:
	slicePrint(headerSlice)
	slicePrint(maaltijdSlice)
}

// Declaring the slice printing function:
func slicePrint(inputSlice []string) {
	// Iterate over inputSlice and print the output and a tab after:
	for i, output := range inputSlice {
		if i%3 == 0 && i != 0 {
			fmt.Printf("\n")
		}
		fmt.Printf(output + "\t")
		// If i is the same as the length -1, place a newline after:
		if i == (len(inputSlice)-1) && i != 0 {
			fmt.Printf("\n")
		}
	}
}

// Declare the readLine function:
func readLine(query string) string {
	// Print the query and end with a newline:
	fmt.Printf(query + "\n")
	// Create a new scanner which reads from the standard input and scan:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Return the scanned input:
	return scanner.Text()
}

// Define checkLongestOne function:
func checkLongestOne(inputSlice []string, headerSlice []string) {
	biggestOneFirst := 0
	biggestOneSecond := 0
	biggestOneThird := 0
	// Check which entry is the longest:
	for i, output := range inputSlice {
		counter := i + 1
		if counter%1 == 0 && counter%2 != 0 && counter%3 != 0 {
			if len(output) >= biggestOneFirst {
				biggestOneFirst = len(output)
			}
		}
		if counter%2 == 0 && counter%3 != 0 {
			if len(output) >= biggestOneSecond {
				biggestOneSecond = len(output)
			}
		}
		if counter%3 == 0 && counter%2 != 0 {
			if len(output) >= biggestOneThird {
				biggestOneThird = len(output)
			}
		}
	}
	// Add spaces behind the slices until it matches the length of the biggest entry (divided by columns):
	for i, output := range inputSlice {
		counter := i + 1
		var addSpacesHeaders int
		var headerStringSlice string
		if i < len(headerSlice) {
			addSpacesHeaders = biggestOneFirst - len(headerSlice[i])
			headerStringSlice = headerSlice[i]
			for numbers := 0; numbers < addSpacesHeaders; numbers++ {
				headerStringSlice = fmt.Sprint(headerStringSlice + " ")
				headerSlice[i] = headerStringSlice
			}
		}
		if counter%1 == 0 && counter%2 != 0 && counter%3 != 0 {
			if len(output) < biggestOneFirst {
				addSpaces := biggestOneFirst - len(output)

				stringSlice := output
				for numbers := 0; numbers < addSpaces; numbers++ {
					stringSlice = fmt.Sprint(stringSlice + " ")
					inputSlice[i] = stringSlice
				}
			}
		}
		if counter%2 == 0 && counter%3 != 0 {
			if len(output) < biggestOneSecond {
				addspaces := biggestOneSecond - len(output)
				stringSlice := output
				for numbers := 0; numbers < addspaces; numbers++ {
					stringSlice = fmt.Sprint(stringSlice + " ")
					inputSlice[i] = stringSlice
				}
			}
		}
		if counter%3 == 0 && counter%2 != 0 {
			if len(output) < biggestOneThird {
				addspaces := biggestOneThird - len(output)
				stringSlice := output
				for numbers := 0; numbers < addspaces; numbers++ {
					stringSlice = fmt.Sprint(stringSlice + " ")
					inputSlice[i] = stringSlice
				}
			}
		}
	}
}

func userInput() {
	maaltijdSlice = append(maaltijdSlice, readLine("Wat wil jij kopen?"))
	maaltijdSlice = append(maaltijdSlice, readLine("Hoeveel wil jij ervan kopen?"))
	maaltijdSlice = append(maaltijdSlice, readLine("Voor welke maaltijd is het?"))
}
