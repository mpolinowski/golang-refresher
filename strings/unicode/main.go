package main

import (
	"fmt"
	"unicode"
)

func main() {

	const s = "This 'string' is *never* gonna give you up, *never* gonna let you down..."

	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount :=0

	// Count unicode characters:
	for _, ch := range s {
		if unicode.IsPunct(ch) {
			punctCount++
		}
		if unicode.IsLower(ch) {
			lowerCount++
		}
		if unicode.IsUpper(ch) {
			upperCount++
		}
		if unicode.IsSpace(ch) {
			spaceCount++
		}
		if unicode.Is(unicode.Hex_Digit, ch) {
			hexdigitCount++
		}
	}

	fmt.Println("Punctuations", punctCount)
	fmt.Println("Lower Chars", lowerCount)
	fmt.Println("Upper Chars", upperCount)
	fmt.Println("Space Chars", spaceCount)
	fmt.Println("HexDigital Chars", hexdigitCount)
}