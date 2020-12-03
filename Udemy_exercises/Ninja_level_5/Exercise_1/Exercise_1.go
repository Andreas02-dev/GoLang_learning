package main

import (
	"fmt"
)

type person struct {
	firstName        string
	lastName         string
	favoriteIcecream []string
}

func main() {
	p1 := person{
		firstName:        "Andreas",
		lastName:         "Hoornstra",
		favoriteIcecream: []string{"Boba", "Fett"},
	}
	p2 := person{
		firstName:        "Anniek",
		lastName:         "Aarsen",
		favoriteIcecream: []string{"Fett", "Boba"},
	}
	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favoriteIcecream {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n%v\n%v\n", p2.firstName, p2.lastName)
	for _, v := range p2.favoriteIcecream {
		fmt.Printf("%v ", v)
	}
}
