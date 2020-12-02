package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

// Data declaration part:
var a float64

func main() {

}

// programma1 takes 2 float64 numbers, adds them together and prints the result:
func programma1(x, y float64) {
	fmt.Println(x + y)
}

// programma2 takes 2 float64 numbers, adds them together, checks whether they are greater than 30 and returns the boolean result:
func programma2(x, y float64) bool {
	if x+y > 30 {
		return true
	}
	return false
}

// programma3 adds a number to an already declared number:
func programma3(y float64) {
	a = a + y
}

// programma4 reads 10 files with logical filenames, starting with 01:
func programma4(filename, extension string) {
	// Do this 10 times:
	for i := 1; i <= 10; i++ {
		var filestring string
		if i >= 10 {
			// filestring is the filename plus the i value + extension:
			filestring = filename + fmt.Sprint(i) + extension
		} else {
			// filestring is the filename plus 0 plus the i value + extension
			// So that filenames lower than 10 will be with a leading zero:
			filestring = filename + "0" + fmt.Sprint(i) + extension
		}
		// Print the contents of the file using the ioutil.ReadFile function and print errors if any:
		bytes, err := ioutil.ReadFile(filestring)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bytes))
		}
	}
}

// programma5 takes the radius of a circle (r) and returns the circumference of the circle:
func programma5(r float64) float64 {
	return math.Pi * (2 * r)
}
