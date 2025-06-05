package main

import (
	"fmt"

)

func main(){

	const pi float64 = 3.14 // constant declaration
	const greeting string = "Hello, World!" // constant declaration
	const maxLimit int = 100 // constant declaration

	fmt.Println("Value of pi is:", pi) // printing the value of pi
	fmt.Println("Greeting message is:", greeting) // printing the greeting message
	fmt.Println("Maximum limit is:", maxLimit) // printing the maximum limit

	// Uncommenting the following lines will cause a compilation error because constants cannot be reassigned
	// pi = 3.14159 // This will cause an error
	// greeting = "Hi there!" // This will cause an error
	// maxLimit = 200 // This will cause an error

	//how to declair multiple constants in a single line
	const (
		day = "monday"
		month = "january"
		year = 2023 

	)

}