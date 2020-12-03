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
	// Create a person:
	p1 := person{
		firstName:        "Andreas",
		lastName:         "Hoornstra",
		favoriteIcecream: []string{"Boba", "Fett"},
	}
	// Create a person:
	p2 := person{
		firstName:        "Anniek",
		lastName:         "Aarsen",
		favoriteIcecream: []string{"Fett", "Boba"},
	}
	// Print the first and last name of p1:
	fmt.Println(p1.firstName, p1.lastName)
	// Loop over the favorite ice cream slice of p1 and print the values:
	for _, v := range p1.favoriteIcecream {
		fmt.Printf("%v ", v)
	}
	// Print the first and last name of p2:
	fmt.Printf("\n%v\n%v\n", p2.firstName, p2.lastName)
	// Loop over the favorite ice cream slice of p2 and print the values:
	for _, v := range p2.favoriteIcecream {
		fmt.Printf("%v ", v)
	}
}
