package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Data declaration part:

// Shop is the Shop struct:
type Shop struct {
	SalesArray    [6]Sales
	ShopItemArray [3]ShopItem
}

// Sales is the Sales struct:
type Sales struct {
	KlantNaam        string
	ItemNaam         string
	HoeveelheidItems int
}

// ShopItem is the ShopItem struct:
type ShopItem struct {
	ProductNaam    string
	VerkoopPrijs   float64
	InkoopPrijs    float64
	VoorraadAantal int
}

// Items is a slice of ShopItems:
var Items [3]ShopItem = [3]ShopItem{
	{
		ProductNaam:    "Broodje kroket",
		VerkoopPrijs:   1.50,
		InkoopPrijs:    0.50,
		VoorraadAantal: 25},
	{
		ProductNaam:    "Broodje gezond",
		VerkoopPrijs:   1.50,
		InkoopPrijs:    0.50,
		VoorraadAantal: 5},
	{
		ProductNaam:    "Appel",
		VerkoopPrijs:   0.50,
		InkoopPrijs:    0.15,
		VoorraadAantal: 1},
}

// SalesSlice is a slice of Sales:
var SalesSlice [6]Sales = [6]Sales{
	{
		KlantNaam:        "Bob",
		ItemNaam:         "Broodje kroket",
		HoeveelheidItems: 4,
	},
	{
		KlantNaam:        "Gerrit",
		ItemNaam:         "Broodje gezond",
		HoeveelheidItems: 3,
	},
	{
		KlantNaam:        "Björn",
		ItemNaam:         "Appel",
		HoeveelheidItems: 1,
	},
	{
		KlantNaam:        "Student 1",
		ItemNaam:         "Broodje kroket",
		HoeveelheidItems: 5,
	},
	{
		KlantNaam:        "Student 2",
		ItemNaam:         "Broodje kroket",
		HoeveelheidItems: 7,
	},
	{
		KlantNaam:        "Student 3",
		ItemNaam:         "Broodje kroket",
		HoeveelheidItems: 4,
	},
}

// CoolShop is a Shop
var CoolShop Shop = Shop{SalesSlice, Items}

func main() {
	bs, err := json.Marshal(CoolShop)
	if err != nil {
		fmt.Println(err)
		panic(err.Error)
	}
	fmt.Println(string(bs))
	err = ioutil.WriteFile("CHESnackcorner.json", bs, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err.Error)
	}
}
