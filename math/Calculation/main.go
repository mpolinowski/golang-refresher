package main

import (
	"fmt"
	"math"
)

func main() {

	x := 10.0

	// Absolute value
	fmt.Println(math.Abs(x), math.Abs(-x))

	// Power (x^y) and Exp (e^x)
	fmt.Println(math.Pow(x, 2.0))
	fmt.Println(math.Exp(x))

	// Trigonometry
	fmt.Println(math.Cos(math.Pi))
	fmt.Println(math.Sin(2 * math.Pi))
	fmt.Println(math.Tan(math.Pi / 2))

	// Square and cube roots
	fmt.Println(math.Sqrt(144))
	fmt.Println(math.Cbrt(125))

	// Hypothenuse
	fmt.Println(math.Hypot(30, 40))
}