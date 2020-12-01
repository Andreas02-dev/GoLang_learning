package main

import (
	"fmt"
)

func main() {
	// Call the onevenNummers function:
	onevenNummers()
	// Call the evenNummers function:
	evenNummers()
	// Call the deelbaarDoorDrie function:
	deelbaarDoorDrie()
}

// onevenNummers prints all the uneven numbers until 99:
func onevenNummers() {
	for i := 1; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// evenNummers prints all the even numbers until 99:
func evenNummers() {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// deelbaarDoorDrie prints all the numbers until 99 which are divisible by 3:
func deelbaarDoorDrie() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
