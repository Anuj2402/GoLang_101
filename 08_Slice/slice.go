// What are slices in Go?
// A slice is a dynamically-sized, flexible view into the elements of an array.
// Slices are more powerful than arrays and are used frequently in Go programming.
// They provide a way to work with sequences of data without needing to know the size of the data in advance.

package main

import (
	"fmt"
	"slices"
	"sort" // Importing the sort package to work with slices
)

func  main()  {
	
	var numbers []int // Declare a slice of integers
	fmt.Println("Initial slice:", numbers) // Output: Initial slice: []
	fmt.Println("The numbers slice is null", numbers== nil) // Output: The numbers slic is null true
	fmt.Println("Length of the slice:", len(numbers)) // Output: Length of the slice: 0

	// initializing a slize using make function
	nums := make([]int, 0 , 5) // Create a slice with length 0 and capacity 5
	fmt.Println("Slice created using make:", nums) // Output: Slice created using make: []
	fmt.Println("Length of the slice created using make:", len(nums)) // Output: Length of the slice created using make: 0
	fmt.Println("Capacity of the slice created using make:", cap(nums)) // Output: Capacity of the slice created using make: 5
	// Adding elements to a slice using append function
	nums = append(nums, 1,2,3,4,5) // Append elements to the slice
	fmt.Println("Slice after appending elements:", nums) // Output: Slice after appending elements: [1 2 3 4 5]
	fmt.Println("Length of the slice after appending elements:", len(nums)) // Output: Length of the slice after appending elements: 5
	fmt.Println("Capacity of the slice after appending elements:", cap(nums)) // Output: Capacity of the slice after appending elements: 5
	fmt.Println("Slice after appending more than capacity element ", append(nums, 100)) // Output: Slice after appending more than capacity element  [1 2 3 4 5 100]
	fmt.Println("Length of the slice after appending more than capacity element:", len(append(nums, 100))) // Output: Length of the slice after appending more than capacity element: 6
	fmt.Println("Capacity of the slice after appending more than capacity element:", cap(append(nums, 100))) // Output: Capacity of the slice after appending more than capacity element: 10


	// cpying a slice
	slice1 := make([]int, 0, 5) // Create a slice with length 0 and capacity 5
	slice1 = append(slice1, 1, 2, 3, 4, 5) // Append elements to the slice1
	slice2 := make([]int, len(slice1)) // Create a new slice2 with the same length as slice1

	copy(slice2, slice1) // Copy elements from slice1 to slice2
	fmt.Println("Slice1:", slice1) // Output: Slice1: [1 2 3 4 5]
	fmt.Println("Slice2 after copying from Slice1:", slice2) // Output: Slice2 after copying from Slice1: [1 2 3 4 5]

	// Slicing a slice

	slice3 := slice1[1:4] // Create a new slice3 from slice1, taking elements from index 1 to 3 (exclusive)
	fmt.Println("Slice3 (sliced from Slice1):", slice3) // Output: Slice3 (sliced from Slice1): [2 3 4]
	// Slicing a slice with capacity
	slice4 := slice1[1:4:5] // Create a new slice4 from slice1, taking elements from index 1 to 3 (exclusive) with capacity 5
	fmt.Println("Slice4 (sliced from Slice1 with capacity):", slice4) // Output: Slice4 (sliced from Slice1 with capacity): [2 3 4]
	// Modifying a slice
	slice4[0] = 10 // Modify the first element of slice4
	fmt.Println("Slice4 after modification:", slice4) // Output: Slice4 after modification: [10 3 4]
	fmt.Println("Slice1 after modifying Slice4:", slice1) // Output: Slice1 after modifying Slice4: [1 10 3 4 5]
	// Iterating over a slice using a for loop
	fmt.Println("Iterating over Slice1:")
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Element at index %d: %d\n", i, slice1[i])
	}
	// Iterating over a slice using a range loop
	fmt.Println("Iterating over Slice1 using range:")
	for index, value := range slice1 {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}
	// Taking user input to fill a slice
	var userSlice []int
	fmt.Println("Enter 5 integers to fill the slice:")
	for i := 0; i < 5; i++ {
		var input int 
		fmt.Printf("Enter number at index %d: ", i+1)
		fmt.Scanln(&input) // Read user input for each element
		userSlice = append(userSlice, input) // Append the input to the slice
	}

	// Print the user input slice
	fmt.Println("You entered:")
	for i, v := range userSlice {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}

	// String slice
	var names []string // Declare a slice of strings
	names = []string{"Anuj", "Kumar"} // Initialize the slice with values
	fmt.Println("Names slice:", names) // Output: Names slice: [Anuj Kumar]
    fmt.Println("slice the names slice:", names[0:1]) // Output: slice the names slice: [Anuj]

	//packages in slices
	
	// Importing a package to work with slices
	// In Go, you can use the "sort" package to sort slices.

	// Sorting a slice of integers
	sort.Ints(slice1) // Sort the slice1 in ascending order
	fmt.Println("Slice1 after sorting:", slice1) // Output: Slice1 after sorting: [1 3 4 5 10]
	// Sorting a slice of strings
	sort.Strings(names) // Sort the names slice in alphabetical order
	fmt.Println("Names slice after sorting:", names) // Output: Names slice after sorting: [Anuj Kumar]
	// Sorting a slice of strings in reverse order
	sort.Sort(sort.Reverse(sort.StringSlice(names))) // Sort the names slice in reverse alphabetical order
	fmt.Println("Names slice after reverse sorting:", names) // Output: Names slice after reverse sorting: [Kumar Anuj]
	// equality check of slices
	slice5 := []int{1, 2, 3, 4, 5}
	slice6 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slices are queail", slices.Equal(slice5, slice6)) // Output: Slices are queail true
	


}