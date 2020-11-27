package main

import (
	"fmt"
	"math"
)

func main() {
	// Call the arithmeticPower function to multiply x 10 times with itself:
	fmt.Println(arithmeticPower(2))
	// Call the mathPackagePower function to multiply x 10 times with itself:
	fmt.Println(mathPackagePower(2))
	// Round 3.6 to 4:
	fmt.Println(math.Round(3.6))
	// Round to the lowest number, 3.6 will now be 3:
	fmt.Println(math.Floor(3.6))
	// Take the square root of 4, which is 2:
	fmt.Println(math.Sqrt(4))
	// Print the absolute value of -2, which is 2:
	fmt.Println(math.Abs(-2))
	// Print the biggest number of the 2, should be 4:
	fmt.Println(math.Max(2, 4))
	// Print the smallest number of the 2, should be 2:
	fmt.Println(math.Min(2, 4))
	// Print the number of Pi:
	fmt.Println(math.Pi)
}

// arithmeticPower takes a float64 x and returns a float64 y which is x multiplied with itself 10 times:
func arithmeticPower(x float64) (y float64) {
	y = x * x * x * x * x * x * x * x * x * x
	return y
}

// mathPackagePower takes a float64 x and returns a float64 y which is x to the power of 10:
func mathPackagePower(x float64) (y float64) {
	y = math.Pow(x, 10)
	return y
}
