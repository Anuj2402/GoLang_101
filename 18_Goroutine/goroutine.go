/*
	what are goroutines?
	Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions, enabling efficient multitasking within a Go program.
	Goroutines are created using the `go` keyword followed by a function call. They run concurrently with other goroutines and can communicate with each other using channels.
	Goroutines are useful for performing tasks concurrently, such as handling multiple requests in a web server or processing data in parallel, without the overhead of traditional threads.
	Goroutines are designed to be efficient and scalable, allowing developers to write concurrent programs easily without worrying about low-level thread management.

	what are wait groups in Goroutine ?
	Wait groups are a synchronization primitive in Go that allows you to wait for a collection of goroutines to finish executing. They are used to block the main goroutine until all specified goroutines have completed their tasks.
	Wait groups are created using the `sync.WaitGroup` type, which provides methods to add, wait, and mark goroutines as done. They are useful for coordinating the completion of multiple goroutines, ensuring that the main program does not exit before all tasks are completed.
	Wait groups are particularly useful in scenarios where you need to wait for multiple concurrent operations to finish before proceeding, such as when processing multiple requests or performing parallel computations.
	Wait groups are implemented using the `sync` package in Go, which provides a simple and efficient way to manage concurrency and synchronization in your programs.
*/

package main

import (
	"fmt"
	"time"
)
// Function to simulate a task that takes time to complete
func doTask(taskName string, duration time.Duration) {
	fmt.Printf("Starting task: %s\n", taskName) // Print the name of the task being started
	time.Sleep(duration) // Simulate a delay to represent the time taken by the task
	fmt.Printf("Completed task: %s\n", taskName) // Print the name of the task that has been completed
}
func main() {
	// Start a goroutine to perform a task
	go doTask("Task 1", 2*time.Second) // This will run concurrently with the main function

	// Start another goroutine to perform a different task
	go doTask("Task 2", 3*time.Second) // This will also run concurrently with the main function

	for i := 0; i < 5; i++ {
		fmt.Println("Main function is running...") // Print a message indicating that the main function is still running

		go func(i int){ // Start a goroutine that prints a message
			time.Sleep(1 * time.Second) // Simulate a delay to allow the goroutine to run concurrently
			fmt.Println("Goroutine is running...", i) // Print a message from the goroutine

		}(i) // Pass the loop variable i to the goroutine to avoid closure issues
	}

	// Wait for a while to allow the goroutines to complete before the main function exits
	time.Sleep(5 * time.Second) // Wait for 5 seconds to ensure both tasks are completed
	fmt.Println("All tasks completed") // Print a message indicating that all tasks are done
}


