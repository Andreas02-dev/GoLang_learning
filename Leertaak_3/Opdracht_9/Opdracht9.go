package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Data declaration part:
var (
	gebruikers         map[string]string
	ok                 bool
	incorrectTries     int
	correctCredentials bool
	bestelArray        [3]string
	repeat             bool = true
)

const (
	allowedTries int = 3
)

func main() {
	// Initialize standard values for the gebruikers map:
	gebruikers = map[string]string{"Bezoeker1": "12345", "Bezoeker2": "asdfg", "Bezoeker3": "hjkl"}
	// If correct credentials have been entered:
	if correctCredentials, correctGebruikersnaam := checkCredentials(); correctCredentials {
		// Launch the bestelApplicatie:
		bestelApplicatie(correctGebruikersnaam)
	}

}

// readLine takes a query, prints it, scans the input given by the user and returns the string:
func readLine(query string) string {
	// Print the query and end with a newline:
	fmt.Printf(query + "\n")
	// Create a new scanner which reads from the standard input and scan:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Return the scanned input:
	return scanner.Text()
}

func checkCredentials() (bool, string) {
	// Have the credentials be checked for allowedTries times:
	for i := 1; i <= allowedTries; i++ {
		// Ask for the username:
		gebruikersnaamInput := readLine("Voer uw gebruikernaam in:")
		// Ask for the password:
		wachtwoordInput := readLine("Voer uw wachtwoord in:")
		// Check if the username exists and whether the password matches the username-password combination:
		if wachtwoord, ok := gebruikers[gebruikersnaamInput]; ok && wachtwoordInput == wachtwoord {
			// Print welcome with the username:
			fmt.Printf("Welcome, %v!\n", gebruikersnaamInput)
			// Return true:
			return true, gebruikersnaamInput
		}
		// This will execute if the username-password combination does not match the input:
		// Increment incorrectTries to display the amount of tries left for the user:
		incorrectTries++
		fmt.Printf("Username or password incorrect.\n%v tries remaining\n", 3-incorrectTries)
	}
	// Return false after 3 incorrect tries:
	return false, ""
}

// bestelApplicatie launches the main CLI:
func bestelApplicatie(gebruikersnaam string) {
	// Build the bestelArray:
	bestelArray := [3]string{"( 1.) Margherita, €8.50, Tomaten en kaas", "( 2.) Siciliano, €9.25, Kaas, tomaat en pesto", "( 3.) Hawaï, €9.50, Tomaten, kaas, ham en ananas"}
	// Build the filename from the username:
	fileName := gebruikersnaam + ".txt"
	for repeat {
		// Clear the terminal
		clearTerminal()
		// Ask whether customer wants to order or check the order:
		fmt.Println("Pizza Bestelapplicatie - Keuzemenu")
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("( 1.) Nieuwe bestelling aanmaken")
		fmt.Println("( 2.) Bestelling inzien")
		fmt.Println("( 3.) Bestelapplicatie afsluiten")
		fmt.Println("-----------------------------------------------------------------------")
		choiceNumberString := readLine("Wat wilt u doen?")
		clearTerminal()
		choiceNumber, _ := strconv.Atoi(choiceNumberString)
		switch choiceNumber {
		case 1:
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
				orderDetails := fmt.Sprintf("Order: %v\nAantal: %v\nNaam: %v\nWanneer: %v\nOpmerkingen: %v", bestelArray[s], aantal, naam, wanneer, opmerkingen)
				schrijfOrder(fileName, orderDetails)
			}
		case 2:
			if _, err := ioutil.ReadFile(fileName); err != nil {
				fmt.Println("U heeft nog geen bestellingen geplaatst.")
				fmt.Println("-----------------------------------------------------------------------")
				fmt.Println("Press enter to continue")
				readLine("")
			} else {
				fmt.Println("-----------------------------------------------------------------------")
				openOrder(fileName)
				fmt.Println("-----------------------------------------------------------------------")
				fmt.Println("Press enter to continue")
				readLine("")
			}

		case 3:
			repeat = false
		}
	}

}

// clearTerminal clears the terminal
func clearTerminal() {
	for i := 0; i <= 20; i++ {
		fmt.Println("\v")
	}
}

// schrijfOrder takes the filename and the orderDetails strings and writes the details to the file.
func schrijfOrder(filename, orderDetails string) {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("An error has occurred opening the file")
	}
	fullFileContents := string(fileContents) + "\n\n" + orderDetails
	err = ioutil.WriteFile(filename, []byte(fullFileContents), 0655)
	if err != nil {
		fmt.Println("An error has occurred writing to the file")
	}
}

// openOrder takes the filename and prints the contents of the file:
func openOrder(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
