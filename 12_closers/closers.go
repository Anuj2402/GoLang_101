/*
	what are closers?
	Closers are functions that capture the environment in which they were created.
	They can access variables from their surrounding scope even after that scope has exited.
	Closers are useful for creating functions that maintain state or for creating callbacks that can access variables from their surrounding context.
	
	how to make a struct a Private and Public field?
In Go, the visibility (public or private) of a struct or its fields is determined by the case of the first letter:
Public (exported): If the struct or field name starts with an uppercase letter, it is accessible from other packages.
Private (unexported): If the struct or field name starts with a lowercase letter, it is only accessible within the same package.
*/

package main

import (
	"fmt"
)


// Function that returns a closure
// The closure maintains its own state (the count variable)
// Each time the closure is called, it increments the count and returns the new value.

func makeCounter() func() int { //makeCounter is a function that returns a closer func() and that function is retuning an int 
	count := 0 // This variable is captured by the closure
	return func() int { //
		count++ // Increment the count
		return count // Return the current count
	}
}

func main(){
	counter := makeCounter() // Create a new counter closure

	// Call the counter multiple times
	fmt.Println("Counter:", counter()) // Output: Counter: 1
	fmt.Println("Counter:", counter()) // Output: Counter: 2
	fmt.Println("Counter:", counter()) // Output: Counter: 3

	// Each call to the closure maintains its own state
	newCounter := makeCounter() // Create a new counter closure
	fmt.Println("New Counter:", newCounter()) // Output: New Counter: 1
	fmt.Println("New Counter:", newCounter()) // Output: New Counter: 2

}