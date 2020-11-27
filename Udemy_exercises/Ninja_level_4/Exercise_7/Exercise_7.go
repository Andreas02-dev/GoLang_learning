package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}

	for i, v := range xy {
		fmt.Println("Record: ", i)
		for ii, vv := range v {
			fmt.Printf("\t Index: %v \t Value: %v \n", ii, vv)
		}
	}
}
