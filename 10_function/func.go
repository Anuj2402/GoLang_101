/*
	what are fucntions in go?
	Functions in Go are blocks of code that perform a specific task. They can take inputs (parameters) and return outputs (results). Functions help in organizing code, making it reusable and easier to read.
	Functions can be defined with or without parameters and can return multiple values. They can also be anonymous (lambda functions) or named functions.
	Functions can be used to encapsulate logic, making it easier to maintain and test code. They can also be used to create higher-order functions, which take other functions as arguments or return functions as results.
*/

package main

import (
	"fmt"

)
// Defining a simple function
func sayHello() {
	fmt.Println("Hello, World!") // This function prints a greeting message
}

// Function with parameters
func addNumbers(a int, b int) int {
	return a + b // This function takes two integers as parameters and returns their sum

}
// Function to add three numbers
func addThreeNumbers(a, b, c int ) int{
	return a + b + c // This function takes three integers as parameters and returns their sum
}

// Function to perform an operation on two numbers
func performOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b) // This function takes two integers and a function as parameters, and returns the result of applying the function to the integers
}


func main(){

	// Defining a simple function
	sayHello() // Calling the function
	// Output: Hello, World!
	// Function with parameters

	fmt.Println("Enter the first number:")
	var a int
	fmt.Scanln(&a) // Taking user input for the first number
	fmt.Println("Enter the second number:")
	var b int
	fmt.Scanln(&b) // Taking user input for the second number
	fmt.Println(("Enter the Third number:"))
	var c int
	fmt.Scanln(&c) // Taking user input for the third number

	
// Function to add two numbers

	// The function takes two integers as parameters and returns their sum
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, addNumbers(a, b)) // Calling the function and printing the result

	fmt.Printf("The sum of %d, %d and %d is: %d\n", a, b, c, addThreeNumbers(a, b, c)) // Calling the function and printing the result
	// Output: The sum of 5 and 10 is: 15
	

	// passing a function as an argument to another function
	result := performOperation(a, b, addNumbers) // Calling the performOperation function with addNumbers as an argument
	fmt.Printf("The result of performing the operation on %d and %d is: %d\n", a, b, result) // Printing the result
}

