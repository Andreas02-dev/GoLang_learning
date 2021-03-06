package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var rand1 *rand.Rand
var boolWin bool
var intTime int

func main() {
	fmt.Println("Druk op enter als u wilt dobbelen.")
	// Wait until user hits enter
	readLine()
	// Do this 4 times:
	for i := 1; i <= 5; i++ {
		// If boolWin is true, exit:
		if boolWin == true {
			break
		}
		// Generate a new rand entity using generateSource.
		// Variate between 6 source times:
		switch intTime {
		case 0:
			rand1 = rand.New(generateSource1())
		case 1:
			rand1 = rand.New(generateSource2())
		case 2:
			rand1 = rand.New(generateSource3())
		case 3:
			rand1 = rand.New(generateSource4())
		case 4:
			rand1 = rand.New(generateSource5())
		case 5:
			rand1 = rand.New(generateSource6())
		}
		// Generate a random number and print it:
		randomNum := rand1.Intn(6) + 1
		fmt.Printf("U gooit: %v\n", randomNum)
		// Print "U wint!" if number is 4 and break if true:
		switch randomNum {
		case 4:
			fmt.Println("U wint!")
			boolWin = true
		}
		// If i is 5 and no 4 has been given, print the following:
		if i == 5 && boolWin == false {
			fmt.Println("U heeft geen 4 gegooid en u verliest..")
		}
	}
}

// generateSource1 generates a new source for rand using the current time:
func generateSource1() rand.Source {
	intTime = 1
	return rand.NewSource(time.Now().UnixNano())
}

// generateSource2 generates a new source for rand using the current time + 1 hour:
func generateSource2() rand.Source {
	intTime = 2
	return rand.NewSource(time.Now().Add(time.Hour * 1).UnixNano())
}

// generateSource3 generates a new source for rand using the current time + 2 hours:
func generateSource3() rand.Source {
	intTime = 3
	return rand.NewSource(time.Now().Add(time.Hour * 2).UnixNano())
}

// generateSource4 generates a new source for rand using the current time + 3 hours:
func generateSource4() rand.Source {
	intTime = 4
	return rand.NewSource(time.Now().Add(time.Hour * 3).UnixNano())
}

// generateSource5 generates a new source for rand using the current time + 4 hours:
func generateSource5() rand.Source {
	intTime = 5
	return rand.NewSource(time.Now().Add(time.Hour * 4).UnixNano())
}

// generateSource6 generates a new source for rand using the current time + 5 hours:
func generateSource6() rand.Source {
	intTime = 0
	return rand.NewSource(time.Now().Add(time.Hour * 5).UnixNano())
}

// readLine returns the scanned input as a string:
func readLine() string {
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	return scanner1.Text()
}
