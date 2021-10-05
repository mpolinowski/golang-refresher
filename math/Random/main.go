package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Random seed as a starting point for the generator
	rand.Seed(time.Now().UnixNano())

	// Generate random integers
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(5))

	// Generate random floats
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	// Generate random ints between a and b
	const a, b = 1, 10
	for i := 0; i<5; i++ {
		n := a + rand.Intn(b-a+1)
		fmt.Print(n, " " )
	}
	fmt.Println()

	// Generate random uppercase character
	for i := 0; i<5; i++ {
		c := string('A' + rune(rand.Intn('Z'-'A'+1)))
		fmt.Printf("%s", c)
	}
	fmt.Println()


	// Permutations - randomly select from array
	s := []string{"Tomcat", "Fulcrum", "Viper", "Frogfoot", "Apache", "Hind", "Bone"}
	// Reorder slice randomly
	indexes := rand.Perm(len(s))
	fmt.Println(indexes)
	// Print index in the order that was generated
	for i := 0; i<len(indexes); i++ {
		fmt.Println(s[indexes[i]])
	}
 
}