package main

import (
	"fmt"
	"strings"
)

func main() {
	
	var sb strings.Builder

	// Provide content
	sb.WriteString("This string is never gonna give you up \n")
	sb.WriteString("Never gonna let you down \n")
	sb.WriteString("Never gonna say goodbye \n")
	sb.WriteString("Never gonna tell a lie and hurt you \n")

	// Concatenate string
	fmt.Println(sb.String())

	// Get length of build string
	fmt.Println("Builder size:", sb.Len())

	// Builder capacity
	fmt.Println("Capacity:", sb.Cap())

	// Add capacity of 1k to reserve space
	sb.Grow(1024)
	fmt.Println("Capacity:", sb.Cap())

	// Reset content of builder
	sb.Reset()
	fmt.Println("After reset")
	fmt.Println("Capacity:", sb.Cap())
	fmt.Println("Builder Size:", sb.Len())
}