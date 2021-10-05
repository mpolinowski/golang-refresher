package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// Random seed as a starting point for the generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle
	const numstring = "Tomcat Fulcrum Viper Frogfoot Apache Hind Bone"
	// split into array of words
	words := strings.Fields(numstring)
	// generate new order
	rand.Shuffle(len(words), func(i,j int) {
		// make words swap places
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
 

	// Random Strings
	const uppercase = "QWERTYUIOOPASDFGHJKLZXCVBNM"
	const lowercase = "qwertyuiopasdfghjklzxcvbnm"
	const digits = "1234567890"
	const special = "!@#$%^&*()_+<>?:;/.|,']"
	const allchars = uppercase + lowercase + digits + special

	// Generate a 10 digit password
	var sb strings.Builder
	length := 10

	for i := 0; i < length; i++ {
		sb.WriteRune(rune(allchars[rand.Intn(len(allchars))]))
	}
	fmt.Println("String result: ", sb.String())

	// Generate a 10 digit password that !must contain all types of characters
	// Create a buffer that takes 1 random character from all sets
	buf := make([]byte, length)
	buf[0] = uppercase[rand.Intn(len(uppercase))]
	buf[1] = lowercase[rand.Intn(len(lowercase))]
	buf[2] = digits[rand.Intn(len(digits))]
	buf[3] = special[rand.Intn(len(special))]
	// Fill up the rest of the buffer with random `allchars`
	for i := 4; i < length; i++ {
		buf[i] = allchars[rand.Intn(len(allchars))]
	}
	// Make sure that the position of the first 4 chars is shuffled
	rand.Shuffle(len(buf), func(i, j int) {
		// make characters swap places
		buf[i], buf[j] = buf[j], buf[i]
	})
	fmt.Println("String result: ", string(buf))

}