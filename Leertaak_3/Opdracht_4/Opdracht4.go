package main

import (
	"fmt"
	"io/ioutil"
)

// Data declaration part:
var fileCounter int = 1

func main() {
	// Create the file name:
	filename := fmt.Sprintf("order%v.txt", fileCounter)
	// Call the schrijfOrder function to write the orderDetails to the file:
	schrijfOrder(filename, "Margherita, 1, morgen Hawaï, 2, morgen")
	// Create the file name:
	filename = fmt.Sprintf("order%v.txt", fileCounter)
	// Call the schrijfOrder function to write the orderDetails to the file:
	schrijfOrder(filename, "Hawaï, 1, vandaag")
	// Print the first order details:
	fmt.Println("Order 1:")
	openOrder("order1.txt")
	// Print the second order details:
	fmt.Println("Order 2:")
	openOrder("order2.txt")
}

// schrijfOrder takes the filename and the orderDetails strings and writes the details to the file.
// It also increments the fileCounter variable:
func schrijfOrder(filename, orderDetails string) {
	ioutil.WriteFile(filename, []byte(orderDetails), 0655)
	fileCounter++
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
