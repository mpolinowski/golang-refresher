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

func decodeJSON() {
	// Declare JSON data to decode
	data := []byte(`
		{
			"char": "Joe Miller",
			"actor": "Thomas Jane",
			"loc": "Ceres",
			"s": [1,2,3,4]
		}
	`)

	// Declare person struct of type person for the JSON data
	var p person
	// Check if JSON is valid and un-marshal
	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}

}

func mapJSON() {

	// Declare JSON data to decode
	dataMap := []byte(`
		{
			"char": "Camina Drummer",
			"actor": "Cara Gee",
			"loc": "Tycho Station",
			"s": [4,5,6]
		}
	`)
	// Decode JSON into a map structure
	var m map[string]interface{}
	json.Unmarshal(dataMap, &m)
	fmt.Printf("%#v\n", m)

	// Iterate over map entries
	for k,v := range m {
		fmt.Printf("key (%v), value (%T : %v)\n", k, v, v)
	}
}

func main() {
	
	decodeJSON()
	mapJSON()
	
}