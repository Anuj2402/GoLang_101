/*
	what are enums?
	Enums, or enumerations, are a way to define a set of named constants in a programming language.
	In Go, enums are typically implemented using a custom type and a set of constants. They provide a way to group related values together, making the code more readable and maintainable. Enums can be used to represent a fixed set of options or states, such as days of the week, colors, or status codes.
	Enums are not a built-in feature in Go, but they can be created using the `const` keyword and a custom type.
	Enums are useful for improving code clarity and preventing errors by restricting the values that a variable can take to a predefined set of constants.
*/


package main

import (
	"fmt"

)
type orderStatus string // Create a custom type named orderStatus, which is an alias for string
// Define constants for different order statuses
const (
	Pending    orderStatus = "Pending"    // Order is pending
	Processing orderStatus = "Processing" // Order is being processed
	Shipped    orderStatus = "Shipped"    // Order has been shipped
	Delivered  orderStatus = "Delivered"  // Order has been delivered
	Cancelled  orderStatus = "Cancelled"  // Order has been cancelled
	Returned   orderStatus = "Returned"   // Order has been returned
	Refunded   orderStatus = "Refunded"   // Order has been refunded
	Unknown    orderStatus = "Unknown"    // Order status is unknown


)

// Function to print the order status
func printOrderStatus(status orderStatus) {
	fmt.Printf("Order Status: %s\n", status) // Print the order status
}

func main() {
	// Define an enum type for days of the week
	type Day int // Create a custom type named Day, which is an alias for int

	// Define constants for each day of the week
	const (
		Sunday Day = iota // 0  // iota is used to create a sequence of constants starting from 0
		Monday             // 1
		Tuesday            // 2
		Wednesday          // 3
		Thursday           // 4
		Friday             // 5
		Saturday           // 6
	)

	// Print the values of the enum constants
	fmt.Println("Days of the week:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

	fmt.Println("Printing enum values from a map")
	// Another way to use enums is to create a map for better readability
	daysOfWeek := map[Day]string{
		Sunday:    "Sunday",
		Monday:    "Monday",
		Tuesday:   "Tuesday",
		Wednesday: "Wednesday",
		Thursday:  "Thursday",
		Friday:    "Friday",
		Saturday:  "Saturday",

}

	// Print the days of the week using the map
	for day , name := range daysOfWeek {
		fmt.Printf("%s: %d\n", name, day)
	}

	fmt.Println("Printing enum values from a slice")
	// best way to use enums is to create a slice for better readability
	namesSlice := []string{
		"Anuj",
		"Ankit",
		"Anshul",
		"Ankush",
		"Ankita",
		"Anjali",
		"Anamika",

	}
	// Print the names using the slice
	for i, name := range namesSlice {
		fmt.Printf("Name %d: %s\n", i+1, name)
	}

	//call the function to print the order status
	printOrderStatus(Pending)    // Print the status of a pending order
	printOrderStatus(Processing) // Print the status of a processing order
	printOrderStatus(Shipped)    // Print the status of a shipped order
	printOrderStatus(Delivered)  // Print the status of a delivered order
	printOrderStatus(Cancelled)  // Print the status of a cancelled order
	printOrderStatus(Returned)   // Print the status of a returned order
	printOrderStatus(Refunded)   // Print the status of a refunded order
	printOrderStatus(Unknown)    // Print the status of an unknown order
	
}
// Output:
// Days of the week:
// Sunday: 0
// Monday: 1
// Tuesday: 2
// Wednesday: 3		