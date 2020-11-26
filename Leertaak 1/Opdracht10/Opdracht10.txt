package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Data declaration part:
var bestelArray [3]string

func main() {
	// Build the bestelArray:
	bestelArray := [3]string{"( 1.) Margherita, €8.50, Tomaten en kaas", "( 2.) Siciliano, €9.25, Kaas, tomaat en pesto", "( 3.) Hawaï, €9.50, Tomaten, kaas, ham en ananas"}
	// Clear the terminal
	clearTerminal()
	// Print the application name
	fmt.Println("Pizza Bestelapplicatie - Bestellen")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("( 1.) Margherita, €8.50, Tomaten en kaas")
	fmt.Println("( 2.) Siciliano, €9.25, Kaas, tomaat en pesto")
	fmt.Println("( 3.) Hawaï, €9.50, Tomaten, kaas, ham en ananas")
	fmt.Println("-----------------------------------------------------------------------")
	orderNummer := readLine("Wat is het ordernummer van de pizza die u wilt bestellen?")
	aantal := readLine("Hoeveel pizza's wilt u bestellen?")
	naam := readLine("Wat is uw naam?")
	wanneer := readLine("Wanneer wilt u de pizza ontvangen?")
	opmerkingen := readLine("Heeft u nog opmerkingen?")
	clearTerminal()
	fmt.Println("Pizza Bestelapplicatie - Bestellen")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("Ordernummer: " + orderNummer)
	fmt.Println("Aantal: " + aantal)
	fmt.Println("Naam: " + naam)
	fmt.Println("Wanneer: " + wanneer)
	fmt.Println("Opmerkingen: " + opmerkingen)
	// Put a white line
	fmt.Println("-----------------------------------------------------------------------")
	s, _ := strconv.Atoi(orderNummer)
	s--
	if s < 0 {
		fmt.Println("Please enter a valid order number.")
	} else {
		fmt.Println("Order: " + bestelArray[s])
		fmt.Println("Aantal: " + aantal)
		fmt.Println("Naam: " + naam)
		fmt.Println("Wanneer: " + wanneer)
		fmt.Println("Opmerkingen: " + opmerkingen)
		fmt.Println(" ")
		fmt.Println("Fijn, een vaste klant. Uw pizza is onderweg!")
	}

}

// This function clears the terminal
func clearTerminal() {
	for i := 0; i <= 20; i++ {
		fmt.Println("\v")
	}
}

func readLine(query string) string {
	// Print the query and end with a newline:
	fmt.Printf(query + "\n")
	// Create a new scanner which reads from the standard input and scan:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Return the scanned input:
	return scanner.Text()
}
