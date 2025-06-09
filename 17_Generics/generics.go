/*
	what is generics?
	Generics is a feature that allows you to write functions and data structures that can operate on different types without sacrificing type safety.
	Generics enable you to create reusable code that can work with any data type, making your code more flexible and reducing duplication.
	Generics are defined using type parameters, which are placeholders for types that can be specified when the function or data structure is used.
// 	Generics are useful for creating functions and data structures that can work with multiple types, improving code reusability and reducing the need for type-specific implementations.
*/

package main
import (
	"fmt"

)

// Define a generic function that takes a slice of any type and returns the first element
func firstElement[T any](slice []T) T {
	if len(slice) == 0 {
		var zeroValue T // Create a zero value of type T
		return zeroValue // Return the zero value if the slice is empty
	}
	return slice[0] // Return the first element of the slice
}
// Before generics, you would have to write separate functions for each type, like this:
// func firstInt(slice []int) int {
// 	if len(slice) == 0 {
// 		return 0 // Return zero if the slice is empty
// 	}
// 	return slice[0] // Return the first element of the slice
// }
// func firstString(slice []string) string {
// 	if len(slice) == 0 {
// 		return "" // Return empty string if the slice is empty
// 	}
// 	return slice[0] // Return the first element of the slice
// }
// This generic function allows you to handle any type of slice without needing to write separate functions for each type.
// The type parameter T can be any type, and the function will work with slices of that type.
// The `any` constraint means that T can be any type, making the function flexible and reusable for different data types.
// The function checks if the slice is empty and returns a zero value of type T if it is, otherwise it returns the first element of the slice.
// This way, you can use the same function for slices of different types, such as int, string, etc.
// This eliminates the need for multiple functions for different types, making the code cleaner and more maintainable.

//using generics in structs
// You can also use generics in structs to create data structures that can hold values of any type.
 type stack[T any] struct {
	data []T // Slice to hold the elements of the stack
}
// Method to push an element onto the stack
func (s *stack[T]) push (value T){
	s.data = append(s.data, value) // Append the value to the stack

}

func main() {
	// Create a slice of integers
	intSlice := []int{1, 2, 3, 4, 5}
	// Call the generic function with the integer slice
	firstInt := firstElement(intSlice)
	fmt.Println("First element of intSlice:", firstInt) // Output: First element of intSlice: 1

	// Create a slice of strings
	stringSlice := []string{"apple", "banana", "cherry"}
	// Call the generic function with the string slice
	firstString := firstElement(stringSlice)
	fmt.Println("First element of stringSlice:", firstString) // Output: First element of stringSlice: apple

	//Before generics, you would have to write separate functions for each type, like this:
	// firstInt := firstInt(intSlice)
	// fmt.Println("First element of intSlice:", firstInt) // Output: First element of intSlice: 1
	// firstString := firstString(stringSlice)
	// fmt.Println("First element of stringSlice:", firstString) // Output: First element of stringSlice: apple
	// You can also use the generic function with other types, such as float64, structs, etc.
	floatSlice := []float64{1.1, 2.2, 3.3}
	firstFloat := firstElement(floatSlice)
	fmt.Println("First element of floatSlice:", firstFloat) // Output: First element of floatSlice: 1.1

	// Using generics in structs
	myStack := stack[int]{} // Create a stack of integers
	myStack.push(10) // Push an integer onto the stack
	fmt.Println("Stack after pushing 10:", myStack.data) // Output: Stack after pushing 10: [10]
	myStack.push(20) // Push another integer onto the stack
	fmt.Println("Stack after pushing 20:", myStack.data) // Output: Stack after pushing 20: [10 20]
	myStringStack := stack[string]{} // Create a stack of strings
	myStringStack.push("hello") // Push a string onto the stack
	fmt.Println("String stack after pushing 'hello':", myStringStack.data) // Output: String stack after pushing 'hello': [hello]



}

