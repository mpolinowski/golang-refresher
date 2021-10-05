package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {

	x := 26
	f := 337.99

	// Formatting
	// decimal
	fmt.Printf("%d\n", x)
	// hexadecimal
	fmt.Printf("%x\n", x)

	// Booleans
	fmt.Printf("%t\n", x > 10)

	// Floats
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

	// Explicit argument indexes
	fmt.Printf("%d %d\n", 69, 119)
	fmt.Printf("%[2]d %[1]d\n", 69, 119)

	// Explicit argument indexes with repeated values
	fmt.Printf("%d %d %#[1]o %#[2]x\n", 69, 119)

	// Print in default format
	c := circle {
		radius: 456,
		border: 2,
	}

	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)

	// Use SprintF to return value as string
	s := fmt.Sprintf("%d %d", 69, 119)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

}