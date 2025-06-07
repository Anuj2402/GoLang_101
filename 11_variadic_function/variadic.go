/*
	what is a variadic function?
	A variadic function is a function that can accept a variable number of arguments.
    Variadic functions are useful when you want to pass a variable number of arguments to a function without having to define each argument explicitly.

*/

package main

import (
	"fmt"

)

// Variadic function to sum a variable number of integers
// what is ...int here? -> // The ...int syntax indicates that the function can accept a variable number of int arguments.
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num // Iterate over the numbers and sum them up
	}
	return total // Return the total sum

}
func main(){

	result := sum(1, 2, 3, 4, 5) // Calling the variadic function with multiple arguments
	fmt.Println("The sum is:", result) // Output: The sum is: 15

	//passing a slice to a variadic function
	nums := []int{10, 20, 30, 40} // Creating a slice of integers
	result = sum(nums...) // Using the ... operator to pass the slice as individual arguments
	fmt.Println("The sum of slice is:", result) // Output: The sum of slice is: 100


}

