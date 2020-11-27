package main

import (
	"fmt"
)

func main() {
	// Notice how the exercise explicitly states between, while the answer has 100 in it too
	for i := 10; i < 100; i++ {
		fmt.Printf("Remainder for number %v divided by 4: %d\n", i, (i % 4))
	}
}
