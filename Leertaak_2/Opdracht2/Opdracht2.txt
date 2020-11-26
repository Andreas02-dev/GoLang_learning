package main

import (
	"fmt"
)

func main() {
	x := floatScanner("Vul uw eerste variabele in:")
	y := floatScanner("Vul uw tweede variabele in:")
	fmt.Println(x + y)
}

func floatScanner(query string) (x float64) {
	fmt.Println(query)
	fmt.Scanln(&x)
	return x
}
