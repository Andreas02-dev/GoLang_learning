package main

import (
	"fmt"
)

// Data declaration part:
var namesArray [3]string = [3]string{"Olav", "Josefien", "Daniël"}
var gemaakteKostenArray [3]float64 = [3]float64{45.00, 30.00, 10.00}
var aantalKeerGekooktArray [3]int = [3]int{4, 3, 2}
var value float64
var index float64
var valueInt int

func main() {
	// Operatie 1:
	balansKookbeurtAanpassen(8.50, 0)
	// Operatie 2:
	gekooktDecrementeren(1)
	kostenInvoeren(-7.30, 1)
	// Operatie 3:
	for i := 1; i <= 3; i++ {
		balansKookbeurtAanpassen(9.30, 2)
	}
	// Print the names of all the people:
	fmt.Println("Names:")
	for _, v := range namesArray {
		fmt.Println(v)
	}
	// Print the costs per person:
	fmt.Println("Costs:")
	for _, v := range gemaakteKostenArray {
		fmt.Println(v)
	}
	// Print the amount of times a person has cooked:
	fmt.Println("Amount of times cooked:")
	for _, v := range aantalKeerGekooktArray {
		fmt.Println(v)
	}
	// Print the name of the person which has to cook, cost-wise:
	fmt.Println("Has to cook cost-wise:")
	fmt.Println(namesArray[minsteKosten()])
	// Print the name of the person which has to cook given the least amount of cooking times:
	fmt.Println("Has to cook, amount of cook-times wise:")
	fmt.Println(namesArray[minsteKookbeurten()])
	// Print the total costs made:
	gemaakteKosten := 0.00
	for _, v := range gemaakteKostenArray {
		gemaakteKosten = gemaakteKosten + v
	}
	fmt.Printf("Totale gemaakte kosten: %v\n", gemaakteKosten)
	// Print the average costs per cook-time:
	fmt.Println("Gemiddelde kosten per kookbeurt:")
	aantalKookbeurten := 0
	for _, v := range aantalKeerGekooktArray {
		aantalKookbeurten = aantalKookbeurten + v
	}
	fmt.Println(gemaakteKosten / float64(aantalKookbeurten))
	// Print the average costs per cook-time per person:
	fmt.Println("Costs per cooktime per person:")
	fmt.Println(gemaakteKosten / float64(aantalKookbeurten) / float64(len(namesArray)))
}

// kostenInvoeren voert de kosten in voor de persoon:
func kostenInvoeren(kosten float64, index int) {
	gemaakteKostenArray[index] = gemaakteKostenArray[index] + kosten
}

// gekooktIncrementeren incrementeert de aantalKeerGekookt voor de persoon:
func gekooktIncrementeren(index int) {
	aantalKeerGekooktArray[index]++
}

// gekooktDecrementeren decrementeert de aantalKeerGekookt voor de persoon:
func gekooktDecrementeren(index int) {
	aantalKeerGekooktArray[index]--
}

// balansKookbeurtAanpassen voert de kosten in voor de persoon en incrementeert de aantalKeerGekookt voor de persoon:
func balansKookbeurtAanpassen(kosten float64, index int) {
	kostenInvoeren(kosten, index)
	gekooktIncrementeren(index)
}

// minsteKosten returns the index of the person with the smallest costs:
func minsteKosten() (index int) {
	for i, v := range gemaakteKostenArray {
		// If index is 0, set value as the costs of this person:
		if i == 0 {
			value = v
			index = i
		}
		// If the costs of the person is smaller than the smallest value, set costs as value and index as index:
		if v <= value {
			value = v
			index = i
		}
	}
	// Return index after looping over the length of the costs array:
	return index
}

// minsteKosten returns the index of the person with the least amount of times cooked:
func minsteKookbeurten() (index int) {
	for i, v := range aantalKeerGekooktArray {
		// If index is 0, set value as the amount of times cooked of this person:
		if i == 0 {
			valueInt = v
			index = i
		}
		// If the costs of the person is smaller than the smallest value, set amount of times cooked as value and index as index:
		if v <= valueInt {
			valueInt = v
			index = i
		}
	}
	// Return index after looping over the length of the amount of times cooked array:
	return index
}
