package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// The if statement will always execute here:
	if true {
		// Print "ping" to the terminal:
		fmt.Println("ping")
	} else {
		// Print "pong" to the terminal:
		fmt.Println("pong")
	}
	// The else statement will always execute here:
	if false {
		// Print "ping" to the terminal:
		fmt.Println("ping")
	} else {
		// Print "pong" to the terminal:
		fmt.Println("pong")
	}
	// Pass the scanned input into In:
	in := readLine("Please input either 1 or 2")
	// Convert the string to an int:
	input, _ := strconv.Atoi(in)
	if input == 1 {
		// Print "ping" to the terminal:
		fmt.Println("ping")
	} else {
		// Print "pong" to the terminal:
		fmt.Println("pong")
	}

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
