package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Call the readLine function and print the returned input:
	fmt.Println("Gegroet, " + readLine("Wat is uw voornaam?"))
}

// Declare the readLine function:
func readLine(query string) string {
	// Print the query and end with a newline:
	fmt.Printf(query + "\n")
	// Create a new scanner which reads from the standard input and scan:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Return the scanned input:
	return scanner.Text()
}
