package main

import "fmt"

func main() {

	f := 997.123

	// Decimal precision (e.g. 2 decimal points)
	fmt.Printf("%.2f\n", f)

	// Max width (e.g. 15 spaces) and default precision (this will just
	// add empty spaces in front to make lists look pretty)
	fmt.Printf("%15f\n", f)

	// Padding and precision
	fmt.Printf("%15.2f\n", f)

	// Add a character in front, e.g. a `+`
	fmt.Printf("%+15.3f\n", f)

	// Padding with zeros instead of empty spaces
	fmt.Printf("%015.4f\n", f)

}