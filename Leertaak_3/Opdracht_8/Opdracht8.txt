package main

import (
	"fmt"
)

func main() {
	// Data declaration part:
	controletekst := "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur1 := "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur2 := "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Prorgammeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur3 := "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur4 := "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. "
	// Maximaal aantal te behalen punten
	maxAantalPunten := 120
	// Aantal fouten van programmeur 1
	fouten1 := 5
	// Aantal fouten van programmeur 2
	fouten2 := 80
	// Aantal fouten van programmeur 3
	fouten3 := 10
	// Aantal fouten van programmeur 4
	fouten4 := 30
	// Call the controleFunctie for each string:
	controleFunctie(programmeur1, controletekst)
	controleFunctie(programmeur2, controletekst)
	controleFunctie(programmeur3, controletekst)
	controleFunctie(programmeur4, controletekst)
	// Call the geslaagdFunctie to check which programmer passed and failed the test:
	geslaagdFunctie(fouten1, maxAantalPunten, programmeur1, 1)
	geslaagdFunctie(fouten2, maxAantalPunten, programmeur2, 2)
	geslaagdFunctie(fouten3, maxAantalPunten, programmeur3, 3)
	geslaagdFunctie(fouten4, maxAantalPunten, programmeur4, 4)
}

// controleFunctie checks whether the given string is equal to the check string and prints the result:
func controleFunctie(programmeurNummer, controleTekst string) {
	if programmeurNummer == controleTekst {
		fmt.Println("Is klaar met zijn werk.")
	} else {
		fmt.Println("Is niet klaar met zijn werk.")
	}
}

// geslaagdFunctie checks which programmer passed the test and prints the result:
func geslaagdFunctie(vragenFout, maximaalVragen int, programmeurString string, programmeurNummer int) {
	// Calculate the grade:
	cijfer := ((9 * (maximaalVragen - vragenFout)) / maximaalVragen) + 1
	// If the grade is bigger or equal to 6, the test is passed:
	if cijfer >= 6 {
		// Print which programmer passed the test with the corresponding grade:
		fmt.Printf("Programmeur: %v \t Gehaald met een: %v\n", programmeurNummer, cijfer)
	} else {
		// Print which programmer failed the test with the corresponding grade:
		fmt.Printf("Programmeur: %v \t Gezakt met een: %v\n", programmeurNummer, cijfer)
	}
}
