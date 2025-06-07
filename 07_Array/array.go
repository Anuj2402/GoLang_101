package main

// what is an array in Go?
// An array is a collection of elements of the same type, stored in contiguous memory locations.
// Arrays have a fixed size, which is defined at the time of declaration.

import (
	"fmt"
)

func main() {
	// Declaring an array
	var numbers [5]int // An array of 5 integers

	// Initializing an array
	numbers = [5]int{1, 2, 3, 4, 5}

	// Accessing elements of an array
	fmt.Println("First element:", numbers[0])  // Output: First element: 1
	fmt.Println("Second element:", numbers[1]) // Output: Second element: 2

	// Modifying elements of an array
	numbers[0] = 10
	fmt.Println("Modified first element:", numbers[0]) // Output: Modified first element: 10

	// Length of an array
	fmt.Println("Length of the array:", len(numbers)) // Output: Length of the array: 5

	// Iterating over an array using a for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	// Iterating over an array using a range loop
	for index, value := range numbers {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

	//Taking user input to fill an array
	var userNumbers [5]int
	fmt.Println("Enter 5 integers:")
	for i := 0; i < len(userNumbers); i++ {
		fmt.Printf("Enter number at index %d\n", i+1)
		fmt.Scanln(&userNumbers[i]) // Read user input for each element
	}

	// Print the user input array
	fmt.Println("You entered:")
	for i, v := range userNumbers {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}

	// string array
	var name [2]string // An array of 2 strings
	name[0] = "Anuj"
	name[1] = "Kumar"
	fmt.Println("Name array:", name) // Output: Name array: [Anuj Kumar]

	//Note -> if array is empty and the Type is int, // it will be initialized with zero values, i.e., [0 0 0 0 0] for an array of size 5.
	// note -> if array is empty and the Type is string, // it will be initialized with empty  values, i.e., ["", ""] for an array of size 2.
	// note -> if array is empty and the Type is bool, // it will be initialized with false values, i.e., [false false false false false] for an array of size 5.
	

	// 2D Array
	var matrix [2][3]int // A 2D array with 2 rows and 3 columns of type int
	matrix[0] = [3]int{1, 2, 3} // First row
	matrix[1] = [3]int{4, 5, 6} // Second row
	fmt.Println("2D Array (Matrix):")
	for i := 0; i < len(matrix); i++ { // Iterate through each row
		for j := 0; j < len(matrix[i]); j++ { // Iterate through each column in the row

			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println() // New line after each row
	}
}
