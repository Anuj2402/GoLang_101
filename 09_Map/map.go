/*
	What is a map in Go?
	A map is a built-in data type in Go that implements a hash table. It is an unordered collection of key-value pairs, where each key is unique and maps to a value.
	Maps are used to store data in a way that allows for fast retrieval based on keys.

Maps are useful when you need to associate values with unique keys, such as in a dictionary or a database.
Maps are implemented as hash tables, which means they provide average-case constant time complexity for lookups, insertions, and deletions.
*/
package main

import (
	"fmt"
	"maps"
	
)

func main(){

	//creating a map 
	m := make(map[string]int) // Create a map with string keys and int values
	fmt.Println("Initial map:", m) // Output: Initial map: map[]

	// Adding key-value pairs to a map
	m["TWENTY"] = 20
	m["FIFTY"] = 50
	m["HUNDRED"] = 100
	fmt.Println("Map after adding key-value pairs:", m) // Output: Map after adding key-value pairs: map[HUNDRED:100 FIFTY:50 TWENTY:20]

	// Accessing values in a map
	fmt.Println("Value for key 'TWENTY':", m["TWENTY"]) // Output: Value for key 'TWENTY': 20
	// Checking if a key exists in a map
	value, exists := m["FIFTY"]
	if exists {
		fmt.Println("Value for key 'FIFTY':", value) // Output: Value for key 'FIFTY': 50
	} else {
		fmt.Println("Key 'FIFTY' does not exist in the map")
	}

	// Modifying a value in a map
	m["TWENTY"] = 25 // Modify the value for key 'TWENTY'
	fmt.Println("Map after modifying value for key 'TWENTY':", m) // Output: Map after modifying value for key 'TWENTY': map[HUNDRED:100 FIFTY:50 TWENTY:25]
	// Deleting a key-value pair from a map
	delete(m, "HUNDRED") // Delete the key 'HUNDRED' from the map
	fmt.Println("Map after deleting key 'HUNDRED':", m) // Output: Map after deleting key 'HUNDRED': map[FIFTY:50 TWENTY:25]
	// Iterating over a map using a for loop
	fmt.Println("Iterating over the map:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// Length of a map
	fmt.Println("Length of the map:", len(m)) // Output: Length of the map: 2
	
	//creating map with key as string and value also as string
	stringMap := make(map[string]string) // Create a map with string keys and string values
	stringMap["name"] = "Anuj Kumar"
	stringMap["age"] = "25"
	fmt.Println("String map:", stringMap) // Output: String map: map[age:25 name:Anuj Kumar]
	// Accessing values in a string map
	fmt.Println("Value for key 'name':", stringMap["name"]) // Output: Value for key 'name': Anuj Kumar

	// creating a map without the make function
	anotherMap := map[string]int{
		"ONE":   1,
		"TWO":   2,
		"THREE": 3,		
	}
	fmt.Println("Another map:", anotherMap) // Output: Another map: map[ONE:1 THREE:3 TWO:2]
	// Accessing values in another map
	fmt.Println("Value for key 'TWO':", anotherMap["TWO"]) // Output: Value for key 'TWO': 2

	//checking the error when accessing a non-existing key
	ExistingValue, exists := anotherMap["FOUR"]

	if exists {
		fmt.Println("Value for key 'FOUR':", ExistingValue) // Output: Value for key 'FOUR': 0
	} else {
		fmt.Println("Key 'FOUR' does not exist in the map") // Output: Key 'FOUR' does not exist in the map
	}
	
	// checkoing the equality of two maps
	map1 := map[string]int{"A":1, "B":2, "C":3}
	map2 := map[string]int{"A":1, "B":2, "C":3}
	fmt.Println(("Are map1 and map2 quals "), maps.Equal(map1,map2)) // Output: Are map1 and map2 quals  true
}