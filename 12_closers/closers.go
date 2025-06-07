/*
	what are closers?
	Closers are functions that capture the environment in which they were created.
	They can access variables from their surrounding scope even after that scope has exited.
	Closers are useful for creating functions that maintain state or for creating callbacks that can access variables from their surrounding context.	
*/

package main

import (
	"fmt"
)


// Function that returns a closure
func makeCounter() func() int {
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