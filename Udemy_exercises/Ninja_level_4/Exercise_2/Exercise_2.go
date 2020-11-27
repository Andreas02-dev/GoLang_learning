package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range xi {
		fmt.Printf("index: %d\t value: %d\n", i, v)
	}
	fmt.Printf("%T", xi)
}
