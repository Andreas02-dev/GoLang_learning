package main

import (
	"fmt"
)

// Data declaration part:
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// Create truck v1:
	var v1 truck = truck{
		vehicle: vehicle{
			doors: 4,
			color: "White",
		},
		fourWheel: true,
	}
	// Create sedan v2:
	var v2 sedan = sedan{
		vehicle: vehicle{
			doors: 3,
			color: "Black",
		},
		luxury: true,
	}
	// Print each vehicle:
	fmt.Println(v1)
	fmt.Println(v2)
	// Print the color of each vehicle:
	fmt.Println(v1.color)
	fmt.Println(v2.color)
}
