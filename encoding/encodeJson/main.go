package main

import (
	"encoding/json"
	"fmt"
)

// Set types and assign the keys that should be used e.g. Characters -> char
// Omit empty fields and don't use actor names
type person struct {
	Character		string	`json:"char"`
	Actor			string	`json:"-"`
	PlaceOfBirth	string	`json:"loc"`
	Seasons			[]int	`json:"s,omitempty"`
}

func encodeJSON() {
	// Create JSON data
	people := []person {
		{"Joe Miller", "Thomas Jane", "Ceres", []int{1,2,3,4} },
		{"James Holden", "Steven Strait", "Earth", []int{1,2,3,4,5,6} },
		{"Bobbie Draper", "Frankie Adams", "Mars", []int{2,3,4,5,6} },
		{"Camina Drummer", "Cara Gee", "Tycho Station", []int{4,5,6} },
		{"Jean-Luc Picard", "Patrick Stewart", "Earth", nil },
	}

	// Use Marshal to convert data structure to JSON
	result, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)

	// Use MarshalIndent to convert data structure to formatted JSON
	resultFormatted, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resultFormatted)

}

func main() {
	
	encodeJSON()

	
}