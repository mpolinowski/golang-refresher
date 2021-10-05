package main

import "fmt"

func main() {

	// Print a string
	fmt.Print("a string")

	// Print a string with newline
	fmt.Println("a string followed by a newline")

	// Print string with values
	const fortytwo = 42
	const answer = "answer"
	const everything = "everything"
	fmt.Println("The", answer, "to", everything, "is", fortytwo)

	// Print a slice
	items := []int{33, 66, 99, 666}
	length, err := fmt.Println(items)
	fmt.Println(length, err)

}