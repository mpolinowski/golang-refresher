package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	s1 := "This string is never gonna give you up"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "Never gonna give, never gonna give. Give you up!"

	// Split into substrings
	sub1 := strings.Split(s1, "never")
	fmt.Printf("%q\n", sub1)

	// Split string around whitespaces
	result := strings.Fields(s1)
	fmt.Printf("%q\n", result)

	// Split around punctuation
	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	result2 := strings.FieldsFunc(s3, f)
	fmt.Printf("%q\n", result2)

	// Join substrings
	result3 := strings.Join(s2, "-")
	fmt.Println(result3)

	// Replace substring
	rep := strings.NewReplacer(",", " |", ".", " |", "!", " |")
	result4 := rep.Replace(s3)
	fmt.Println(result4)
}