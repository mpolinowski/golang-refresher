package main

import (
	"fmt"
	"math"
)

func main() {

	float := 997.783

	// Print Pi
	fmt.Println(math.Pi)
	fmt.Println(math.E)

	// Floor and Ceiling
	fmt.Println(math.Floor(float))
	fmt.Println(math.Ceil(math.E))

	// Truncate to return integer value
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))

	// Min-Max to determin the biggest/smallest member
	fmt.Println(math.Max(math.Pi, math.Ln2))
	fmt.Println(math.Min(math.Pi, math.Ln2))

	// Modulo - the mod operator is for floats
	fmt.Println(17 % 5)
	fmt.Println(math.Mod(17.0, 5.0))

	// Round and RoundToEven - to closes integer
	fmt.Printf("%.1f\n", math.Round(10.5))
	fmt.Printf("%.1f\n", math.Round(-10.5))

	fmt.Printf("%.1f\n", math.RoundToEven(10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(11.5))

}