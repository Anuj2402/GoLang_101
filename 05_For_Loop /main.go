package main

import (
	"fmt"

)
// for is only construct in Go to create loops

func main() {

	// Basic for loop

		for i := 1; i <= 10; i++ {

				if i == 5{
					fmt.Println("Skipping iteration number 5")
					continue // skip the rest of the loop when i is 5
				}
				fmt.Println("Iteration number:", i)
			}

	// while loop
		i := 1
		for i <= 10 {
			fmt.Println("Iteration number:", i)
			i++
		}
	// infinite loop
		// Uncommenting the following lines will cause an infinite loop
		// for {
		// 	fmt.Println("This is an infinite loop")
		// 	// You can break the loop using a break statement
		// 	if i > 10 {
		// 		break
		// 	}
		// 	i++
		// }
		
	// range loop -> in version 1.18 and later, you can use the range keyword to iterate over slices, arrays, maps, strings, and channels.
		numbers := []int{1, 2, 3, 4, 5}
		for index, value := range numbers {
			fmt.Printf("Index: %d, Value: %d\n", index, value)
		}

		for i := range 11 {
			fmt.Println("Iteration number: ", i)
		}
}
