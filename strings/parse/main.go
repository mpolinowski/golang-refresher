package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "101"
	num := 100


	// Converting int into string
	// The string function does NOT convert an int into a string!
	newstr1 := string(str)
	fmt.Println("The first string is:", newstr1)
	fmt.Printf("%T\n", newstr1)
	newstr2 := string(num)
	fmt.Println("This is not the string I expected:", newstr2)
	fmt.Printf("%T\n", newstr2)
	// To transform the int use:
	newstr3 := strconv.Itoa(num)
	fmt.Println("The second string is:", newstr3)
	fmt.Printf("%T\n", newstr3)


	// Convert string to int
	newstr4, err := strconv.Atoi(str)
	if (err != nil) {
		fmt.Println(err.Error())
	}
	fmt.Println("This is now an integer:", newstr4)
	fmt.Printf("%T\n", newstr4)


	// Parsing strings into different data types
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(u)
	fmt.Printf("%T\n", u)

	// Format converts values into strings
	s1 := strconv.FormatBool(true)
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)
	s3 := strconv.FormatInt(-42, 10)
	fmt.Println(s3)
	fmt.Printf("%T\n", s3)
	s4 := strconv.FormatUint(42, 10)
	fmt.Println(s4)
	fmt.Printf("%T\n", s4)
}