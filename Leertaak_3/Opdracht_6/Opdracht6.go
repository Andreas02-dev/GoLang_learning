package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Print the current date and time using the standard mask:
	fmt.Println(time.Now())
	// Generate a new source using the current time:
	source1 := rand.NewSource(time.Now().UnixNano())
	// Generate a new rand entity with the source:
	rand1 := rand.New(source1)
	// Generate a random number:
	randomNum := rand1.Intn(2)
	// If the number is 0, print "kop", else print "munt":
	if randomNum == 0 {
		fmt.Println("kop")
	} else {
		fmt.Println("munt")
	}
}
