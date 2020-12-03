package main

import (
	"fmt"
)

// Data declaration part:
type person struct {
	firstName        string
	lastName         string
	favoriteIcecream []string
}

func main() {
	// Create a person p1:
	p1 := person{
		firstName:        "Andreas",
		lastName:         "Hoornstra",
		favoriteIcecream: []string{"Boba", "Fett"},
	}
	// Create a person p2:
	p2 := person{
		firstName:        "Anniek",
		lastName:         "Aarsen",
		favoriteIcecream: []string{"Fett", "Boba"},
	}
	// Create a map which stores p1 and p2 with their lastName as key:
	var personMap map[string]person = map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	// Loop over the values in personMap:
	for _, v := range personMap {
		// Print the firstName and lastName of the values in personMap:
		fmt.Printf("%v %v\n", v.firstName, v.lastName)
		// Loop over the favoriteIcecream slice of all the values in personMap and print them:
		for _, v := range v.favoriteIcecream {
			fmt.Printf("%v ", v)
		}
		// Add a newline between the last value from the favoriteIcecream slice and the next firstName:
		fmt.Println()
	}
}
