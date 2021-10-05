package main

import (
	"fmt"
	"strings"
)

func main() {
	// Become a secret agent by encoding your messages
	s:= "This string is never gonna give you up, never gonna let you down..."
	// As a cypher push every character `+2`
	shift := 2

	// Mapping function
	transform := func (r rune) rune {
		switch {

		// if character is uppercase
		case r >= 'A' && r <= 'Z':
			// get ASCII base code of character and add value of `shift`
			value := int('A') + (int(r) - int('A') + shift)
			// if you hit 91 (which equals 'Z')
			if value > 91 {
				// subtract 26 to start from the beginning 'A'
				value -= 26
			// if you hit the lower threshold
			} else if value < 65 {
				// add 26
				value += 26
			}
			return rune(value)
			
		// if character is lowercase
		case r >= 'a' && r <= 'z':
			// get base code of character and add value of `shift`
			value := int('a') + (int(r) - int('a') + shift)
			// if you hit 91 (which equals 'z')
			if value > 122 {
				// subtract 26 to start from the beginning 'a'
				value -= 26
			// if you hit the lower threshold
			} else if value < 97 {
				// add 26
				value += 26
			}
			return rune(value)
		}
		return r
	}

	// Encode message
	encode := strings.Map(transform, s)
	fmt.Println(encode)

	// Encode message
	shift = -shift
	decode := strings.Map(transform, encode)
	fmt.Println(decode)

}