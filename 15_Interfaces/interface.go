/*
	what are interfaces?
	Interfaces are a way to define a contract for types in Go. They allow you to specify methods that a type must implement without dictating how those methods should be implemented. This enables polymorphism, where different types can be treated uniformly based on the methods they implement.
	Interfaces are defined using the `interface` keyword, followed by a list of method signatures.
	Interfaces are useful for creating flexible and reusable code, allowing different types to be used interchangeably as long as they implement the required methods.

	what is the difference between interface and struct?
	Interfaces define a set of methods that a type must implement, while structs are concrete data types that can hold fields and methods. Interfaces allow for polymorphism, enabling different types to be treated uniformly based on the methods they implement. Structs, on the other hand, are used to group related data together and can implement interfaces by defining the required methods.

	Define interface in simple words?
	An interface is a contract that specifies a set of methods that a type must implement. It allows different types to be treated uniformly based on the methods they provide, enabling polymorphism and flexibility in code design.

*/



package main

import (
	"fmt"
)

// Define an interface named Shape
type Shape interface { // this interface is public, so it can be accessed from other packages as it starts with an uppercase letter
	Area() float64 // Method to calculate the area of the shape
	Perimeter() float64 // Method to calculate the perimeter of the shape
}
// Define a struct named Rectangle that implements the Shape interface
type Rectangle struct {
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}
// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // Calculate the area of the rectangle
}
// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height) // Calculate the perimeter of the rectangle
}
// Define a struct named Circle that implements the Shape interface
type Circle struct {
	Radius float64 // Radius of the circle
}
// Implement the Area method for Circle
func (c Circle) Area() float64{
	return 3.14 * c.Radius * c.Radius // Calculate the area of the circle using πr²

}
// Implement the Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius // Calculate the perimeter of the circle using 2πr
}
// Function to print the area and perimeter of a shape
func printShapeInfo(s Shape) { // This function takes a Shape interface as an argument
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter()) // Print the area and perimeter of the shape
}
func main() {
	// Create a Rectangle instance
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Rectangle:")
	printShapeInfo(rect) // Call the function to print rectangle info

	// Create a Circle instance
	circ := Circle{Radius: 7}
	fmt.Println("Circle:")
	printShapeInfo(circ) // Call the function to print circle info
}

// Output:
// Rectangle:
// Area: 50.00, Perimeter: 30.00
// Circle:
// Area: 153.86, Perimeter: 43.96

// In this example, we define an interface `Shape` with two methods: `Area` and `Perimeter`.
// We then create two structs, `Rectangle` and `Circle`, that implement the `Shape` interface by providing their own versions of the `Area` and `Perimeter` methods.
// The `printShapeInfo` function takes a `Shape` interface as an argument and prints the area and perimeter of the shape.
// This allows us to use different shapes interchangeably as long as they implement the `Shape` interface.

