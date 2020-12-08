package main

import (
	"fmt"
)

// Data declaration part:

type informatieStruct struct {
	NaamHuurder   string
	Adres         string
	Postcode      string
	Plaats        string
	Verhuurstatus string
	Prijs         float64
}

var informatieStructSlice []informatieStruct

func main() {
	// Add all the entries:
	addInstance(0, informatieStruct{"Joop Koopgoot", "Hoefweg 3", "6717 LD", "Ede", "verhuurd", 450})
	addInstance(1, informatieStruct{"[null]", "Hoefweg 34", "6717 LD", "Ede", "te huur", 350})
	addInstance(2, informatieStruct{"Piet Paulusma", "Van Hagestraat 809", "6717 DK", "Ede", "verhuurd", 650})
	addInstance(3, informatieStruct{"Merel Fluisterzacht", "Nieuwe Maanderbuurtweg 1174", "6717 BL", "Ede", "onder optie", 625})
	// Change Nieuwe Maanderbuurtweg 1174 to verhuurd:
	changeInstance(3, informatieStruct{"Merel Fluisterzacht", "Nieuwe Maanderbuurtweg 1174", "6717 BL", "Ede", "verhuurd", 625})
	// Delete Hoefweg 34:
	deleteInstance(findIdentifierByAddress("Hoefweg 34"))
	// Change the rent for Joop Koopgoot:
	changeInstance(0, informatieStruct{"Joop Koopgoot", "Hoefweg 3", "6717 LD", "Ede", "verhuurd", (450 * 1.05)})
	// Print the slice:
	fmt.Println(informatieStructSlice)
}

// addInstance takes the Identifier and an informatieStruct and appends it to the informatieStructSlice:
func addInstance(Identifier int, Structure informatieStruct) {
	informatieStructSlice = append(informatieStructSlice, Structure)
}

// changeInstance takes an Identifier and an informatieStruct and changes the Structure to the given structure for the given Identifier:
func changeInstance(Identifier int, Structure informatieStruct) {
	informatieStructSlice[Identifier] = Structure
}

// deleteInstance takes an Identifier and deletes the entry at the given identifier:
func deleteInstance(Identifier int) {
	informatieStructSlice = append(informatieStructSlice[:Identifier], informatieStructSlice[Identifier+1:]...)
}

// findIdentifierByAddress takes an Address and returns the Identifier for the entry with that address:
func findIdentifierByAddress(Address string) int {
	for i, v := range informatieStructSlice {
		if v.Adres == Address {
			return i
		}
	}
	fmt.Println("Error, address not found.")
	return 0
}
