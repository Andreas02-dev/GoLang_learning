package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(2, 3))
}

func sum(x, y int) (z int) {
	z = x + y
	return z
}
