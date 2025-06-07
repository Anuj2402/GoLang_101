/*
	what is a pointer in Go?
	A pointer is a variable that stores the memory address of another variable.
	Pointers are used to reference and manipulate the value of a variable without copying it.
	Pointers are useful for passing large data structures to functions without copying them, and for modifying the original variable's value.

*/

package main

import (
	"fmt"

)
// Function to increment a value using a pointer
func increment(num *int) {
	*num++ // Dereference the pointer and increment the value it points to
	fmt.Println("Incremented value:", *num) // Output: Incremented value: 43
}

// Note: In Go, pointers are not as commonly used as in languages like C or C++,
// because Go has garbage collection and automatic memory management.
// However, they are still useful for certain scenarios, such as modifying values in functions or working with large data structures.

// Pointers are declared using the * operator, and the address of a variable is obtained using the & operator.


// Pointer example in Go

func main(){
	// Declaring a variable
	var num int = 42

	// Declaring a pointer to the variable
	var ptr *int = &num // The & operator gets the memory address of num

	// Accessing the value using the pointer
	fmt.Println("Value of num:", num)         // Output: Value of num: 42
	fmt.Println("Pointer to num:", ptr)       // Output: Pointer to num: 0x...
	fmt.Println("Value at pointer:", *ptr)    // Output: Value at pointer: 42

	// Modifying the value using the pointer
	*ptr = 100
	fmt.Println("Modified value of num:", num) // Output: Modified value of num: 100

	// Pointer to a pointer
	var ptrToPtr **int = &ptr // Pointer to a pointer
	fmt.Println("Pointer to pointer:", ptrToPtr) // Output: Pointer to pointer: 0x...
	fmt.Println("Value at pointer to pointer:", **ptrToPtr) // Output: Value at pointer to pointer: 100
	// Using pointers in functions
	increment(&num) // Passing the address of num to the function
	fmt.Println("Final value of num:", num) // Output: Final value of num: 43

	//pass by reference vs pass by value
	// In Go, when you pass a variable to a function, it is passed by value by default.
	// This means that a copy of the variable is made, and any changes made to the parameter inside the function do not affect the original variable.
	// However, if you want to modify the original variable, you can pass a pointer to the variable instead.
	// This allows the function to access and modify the original variable directly, as it has the memory address of the variable.
	fmt.Println("Original value of num after increment function:", num) // Output: Original value of num after increment function: 43
	fmt.Println("Pointer to num after increment function:", ptr) // Output: Pointer to num after increment function: 0x...
	fmt.Println("Value at pointer after increment function:", *ptr) // Output: Value at pointer after increment function: 43
	fmt.Println("Pointer to pointer after increment function:", ptrToPtr) // Output: Pointer to pointer after increment function: 0x...
	fmt.Println("Value at pointer to pointer after increment function:", **ptrToPtr) // Output: Value at pointer to pointer after increment function: 43
	fmt.Println("Address of num:", &num) // Output: Address of num: 0x...
	fmt.Println("Address of ptr:", &ptr) // Output: Address of ptr: 0x...
	fmt.Println("Address of ptrToPtr:", &ptrToPtr) // Output: Address of ptrToPtr: 0x...


}