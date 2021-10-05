package main

import (
	"fmt"
	"strings"
)

func main() {

	s:= "This string is never gonna give you up, never gonna let you down..."

	// Print length of a string
	fmt.Println(len(s))

	// Iteration
	for _, ch := range s {
		fmt.Print(string(ch), "-")
	}
	fmt.Println()

	// Operators
	fmt.Println("Sapporro" < "Asahi")
	fmt.Println("Tomcat" > "Hornet")
	fmt.Println("rick" == "Rick")
	fmt.Println("rick" != "Rick")

	// Compare
	compare_01 := strings.Compare("Sapporro", "Asahi")
	fmt.Println(compare_01)
	compare_02 := strings.Compare("Sapporro", "Sapporro")
	fmt.Println(compare_02)

	// Unicode case-folding
	fmt.Println(strings.EqualFold("中國", "台灣"))
	fmt.Println(strings.EqualFold("台灣", "台灣"))

	// Character cases
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s)
	fmt.Println(s1, s2, s3)

}