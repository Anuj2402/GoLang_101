/*
	what are structs?
	Structs are a way to group related data together in Go. They allow you to create complex data types that can hold multiple fields of different types.
	Structs are useful for modeling real-world entities and organizing data in a structured way.
	Structs are defined using the `struct` keyword, followed by a list of fields with their types.
	Structs can be used to create custom data types that encapsulate related data and behavior.
	Structs can be used to create complex data structures, such as linked lists, trees, and graphs.
	Structs can be used to create data types that can be passed around in functions and methods, allowing for better organization and encapsulation of data.

	what is struct embedding?
	Struct embedding is a feature in Go that allows you to include one struct within another struct, effectively creating a "has-a" relationship.
	When you embed a struct, the fields and methods of the embedded struct become part of the outer struct, allowing you to access them directly without needing to reference the embedded struct explicitly.
	Struct embedding is useful for creating reusable components and building complex data structures.
	
*/

package main

import (
	"fmt"
	"time"
)

type Order struct {
	OrderID     int
	CustomerName string // Name of the customer who placed the order
	Product     string // Product name
	Quantity    int // Quantity of the product ordered
	Price       float64 // Price of the product
	createdAt  time.Time // Timestamp when the order was created , nanoseconds precision
}

// receiver type is Order
// Methods can be defined on structs to provide behavior associated with the struct.
func (o Order) TotalPrice() float64 {
	return float64(o.Quantity) * o.Price // Calculate total price based on quantity and price
}

func (o Order) String() string {
	return fmt.Sprintf("Order ID: %d, Customer: %s, Product: %s, Quantity: %d, Price: %.2f, Created At: %s",
		o.OrderID, o.CustomerName, o.Product, o.Quantity, o.Price, o.createdAt.Format(time.RFC3339))
}

// changeing the value of the struct
func (o *Order) UpdateQuantity(newQuantity int) {
	o.Quantity = newQuantity // Update the quantity of the order
}


func main(){

	// Creating an instance of the Order struct
	order := Order{
		OrderID:     1,
		CustomerName: "John Doe",
		Product:     "Laptop",
		Quantity:    2,
		Price:       1500.00,
		createdAt:   time.Now(), // Setting the createdAt field to the current time
	}

	// Accessing fields of the struct
	fmt.Println("Order ID:", order.OrderID) // Output: Order ID: 1
	fmt.Println("Customer Name:", order.CustomerName) // Output: Customer Name: John Doe
	fmt.Println("Product:", order.Product) // Output: Product: Laptop
	fmt.Println("Quantity:", order.Quantity) // Output: Quantity: 2
	fmt.Println("Price:", order.Price) // Output: Price: 1500.00
	fmt.Println("Created At:", order.createdAt) // Output: Created At: 2023-10-01 12:34:56.789123456 +0000 UTC m=+0.000000001

	// Modifying a field of the struct
	order.Quantity = 3
	fmt.Println("Updated Quantity:", order.Quantity) // Output: Updated Quantity: 3

	// Structs can also be used as function parameters or return types

	totalPrice := order.TotalPrice() // Calling a method on the struct to calculate total price
	fmt.Println("Total Price:", totalPrice) // Output: Total Price: 4500.00
	// Using the String method to print the order details
	fmt.Println(order) // Output: Order ID: 1, Customer: John Doe, Product: Laptop, Quantity: 2, Price: 1500.00, Created At: 2023-10-01T12:34:56Z

	// Updating the quantity using a method
	order.UpdateQuantity(100) // Calling the method to update the quantity

	fmt.Println("Updated Order:", order) // Output: Updated Order: Order ID: 1, Customer: John Doe, Product: Laptop, Quantity: 100, Price: 1500.00, Created At: 2023-10-01T12:34:56Z
	// Note: In Go, structs are value types, meaning that when you pass a struct to a function or method, a copy of the struct is made.
	// If you want to modify the original struct, you can use a pointer receiver (e.g., `func (o *Order) UpdateQuantity(newQuantity int)`).

	// Example of using a struct from another package
	languages := struct{
		Name string // Name of the programming language
		Year int // Year of creation
	}{
		Name: "Go",
		Year: 2009,		
	}

	fmt.Println("Programming Language:", languages.Name) // Output: Programming Language: Go
	fmt.Println("Year of Creation:", languages.Year) // Output: Year of Creation: 2009
	
	// Example of struct embedding

	// Embedding a struct within another struct
	type Person struct {
		Name    string // Name of the person
		Age     int // Age of the person
		Address struct {
			Street string // Street address
			City   string // City of residence
			State  string // State of residence
			Zip    string // Zip code
		} // Embedded struct for address

}
	// Creating an instance of the Person struct with embedded Address struct
	person := Person{
		Name: "Alice",
		Age:  30,
		
} 
	// Initializing the embedded Address struct
		person.Address.Street = "123 Main St" 
		person.Address.City = "Wonderland"
		person.Address.State = "Fantasy" 
		person.Address.Zip = "12345"

	fmt.Println("Person Name:", person.Name) // Output: Person Name: Alice
	fmt.Println("Person Age:", person.Age) // Output: Person Age: 30
	fmt.Println("Address:", person.Address.Street, person.Address.City, person.Address.State, person.Address.Zip) // Output: Address: 123 Main St Wonderland Fantasy 12345
	// Struct embedding allows you to access the fields of the embedded struct directly
	fmt.Println("City:", person.Address.City) // Output: City: Wonderland
	fmt.Println("State:", person.Address.State) // Output: State: Fantasy
	

}
