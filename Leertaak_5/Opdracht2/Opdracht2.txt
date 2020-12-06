package main

import (
	"fmt"
)

// Data declaration part:

// Sales is the Sales struct:
type Sales struct {
	klantNaam        string
	itemNaam         string
	hoeveelheidItems int
}

// ShopItem is the ShopItem struct:
type ShopItem struct {
	productNaam    string
	verkoopPrijs   float64
	inkoopPrijs    float64
	voorraadAantal int
}

// Items is a slice of ShopItems:
var Items []ShopItem = []ShopItem{
	{
		productNaam:    "Broodje kroket",
		verkoopPrijs:   1.50,
		inkoopPrijs:    0.50,
		voorraadAantal: 25},
	{
		productNaam:    "Broodje gezond",
		verkoopPrijs:   1.50,
		inkoopPrijs:    0.50,
		voorraadAantal: 5},
	{
		productNaam:    "Appel",
		verkoopPrijs:   0.50,
		inkoopPrijs:    0.15,
		voorraadAantal: 1},
}

// SalesSlice is a slice of Sales:
var SalesSlice []Sales = []Sales{
	{
		klantNaam:        "Bob",
		itemNaam:         "Broodje kroket",
		hoeveelheidItems: 4,
	},
	{
		klantNaam:        "Gerrit",
		itemNaam:         "Broodje gezond",
		hoeveelheidItems: 3,
	},
	{
		klantNaam:        "Björn",
		itemNaam:         "Appel",
		hoeveelheidItems: 1,
	},
	{
		klantNaam:        "Student 1",
		itemNaam:         "Broodje kroket",
		hoeveelheidItems: 5,
	},
	{
		klantNaam:        "Student 2",
		itemNaam:         "Broodje kroket",
		hoeveelheidItems: 7,
	},
	{
		klantNaam:        "Student 3",
		itemNaam:         "Broodje kroket",
		hoeveelheidItems: 4,
	},
}

func main() {
	// Print the ShopItem and SalesSlice slices:
	for _, v := range Items {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println()
	fmt.Println()
	for _, v := range SalesSlice {
		fmt.Printf("%+v\n", v)
	}
	// Print how many Broodjes Kroket Bob has had:
	fmt.Println()
	fmt.Printf("%+v\n", SalesSlice[0])
	fmt.Println()
	// Increment this by one:
	SalesSlice[0].hoeveelheidItems++
	// Delete student 3's sale:
	SalesSlice = SalesSlice[:5]
	for _, v := range SalesSlice {
		fmt.Printf("%+v\n", v)
	}
}
