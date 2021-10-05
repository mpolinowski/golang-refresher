package main

import (
	"fmt"
	"strings"
)

func main() {

	s:= "This string is never gonna give you up, never gonna let you down..."
	vowels := "aeiouAEIOU"

	fname := "textfile.txt"
	fname2 := "imagefile.png"

	// Contains a sub-string
	fmt.Println(strings.Contains(s, "gonna"))

	// ContainsAny of these characters
	fmt.Println(strings.ContainsAny(s, "xyzi"))

	// Find offset of first instance of substring
	fmt.Println(strings.Index(s, "give"))
	fmt.Println(strings.Index(s, "rick"))

	// Find offset of first instance of any of a set of characters
	fmt.Println(strings.IndexAny(s, vowels))
	fmt.Println(strings.IndexAny("xyzs", vowels))

	// Check if strings starts/ends with a substring
	fmt.Println(strings.HasSuffix(fname, "txt"))
	fmt.Println(strings.HasPrefix(fname2, "image"))

	// Count non-overlapping instances of a substring
	fmt.Println(strings.Count(s, "gonna"))
	fmt.Println(strings.Count(s, "is"))

}