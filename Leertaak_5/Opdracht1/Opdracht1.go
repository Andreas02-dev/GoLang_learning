package main

import (
	"fmt"
)

// Data declaration part:
var capitalMap map[string]string = map[string]string{"Nederland": "Amsterdam", "Duitsland": "Berlijn", "Ierland": "Dublin", "Noorwegen": "Oslo", "Frankrijk": "Parijs"}
var capitalMap2 map[string][]string = map[string][]string{"België": {"Brussel"}, "Nederland": {"Amsterdam", "Amersfoort"}, "Duitsland": {"Berlijn", "Brussel", "Oslo"}}

func main() {
	// Print the capital of Ireland:
	fmt.Println(capitalMap["Ierland"])
	// Delete France from capitalMap:
	delete(capitalMap, "Frankrijk")
	// Add Belgium to capitalMap:
	capitalMap["België"] = "Brussel"
	// Print the whole map:
	fmt.Println(capitalMap)
	// Add the Netherlands (Zuid-Holland) again with Den Haag:
	capitalMap["Nederland"] = "Den Haag"
	// Add Koninkrijk der Nederlanden with Amsterdam to capitalMap:
	capitalMap["Koninkrijk der Nederlanden"] = "Amsterdam"
	// Print the whole map:
	fmt.Println(capitalMap)
	// Print the length of the map:
	fmt.Println(len(capitalMap))
	// Print the map in which multiple capitals can be put, using slices:
	fmt.Println(capitalMap2)
}
