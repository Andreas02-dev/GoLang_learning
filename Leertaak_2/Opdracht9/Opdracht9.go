package main

import (
	"fmt"
	"strings"
)

func main() {
	voornaam, achternaam := nameSplitter("Andreas Hoornstra")
	fmt.Printf("Voornaam: %s\nAchternaam: %s", voornaam, achternaam)
}

func nameSplitter(fullName string) (voornaam, achternaam string) {
	s := strings.Split(fullName, " ")
	voornaam = s[0]
	achternaam = s[1]
	return voornaam, achternaam
}
