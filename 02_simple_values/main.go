package main

import "fmt"

func main(){
	fmt.Println("This is an integer value")
	fmt.Println(42) // This is an integer value
	fmt.Printf("Type : %T\n", 42) // Type : int

	fmt.Println("This is a float value")
	fmt.Println(3.14) // This is a float value float32

	fmt.Println("This is a boolean value")
	fmt.Println(true) // This is a boolean value

	fmt.Println("This is a string value")
	fmt.Println("Hello, Go!") // This is a string value
	fmt.Printf("Type : %T\n", "Hello, Go!") // Type : string
	fmt.Println("This is a rune value")
	fmt.Println('G') // This is a rune value
	fmt.Println("This is a complex number value")
	fmt.Println(1 + 2i) // This is a complex number value

}