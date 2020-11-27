package main

import (
	"fmt"
)

func main() {
	i := 2002
	for {
		if i > 2020 {
			break
		}
		fmt.Printf("Year: %d\n", i)
		i++
	}
}
