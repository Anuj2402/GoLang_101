/*
	what are wait groups?
	Wait groups are a synchronization primitive in Go that allows you to wait for a collection of goroutines to finish executing. They are used to block the main goroutine until all specified goroutines have completed their tasks.
	Wait groups are created using the `sync.WaitGroup` type, which provides methods to add, wait, and mark goroutines as done. They are useful for coordinating the completion of multiple goroutines, ensuring that the main program does not exit before all tasks are completed.
	Wait groups are particularly useful in scenarios where you need to wait for multiple concurrent operations to finish before proceeding, such as when processing multiple requests or performing parallel computations.
	 Wait groups are implemented using the `sync` package in Go, which provides a simple and efficient way to manage concurrency and synchronization in your programs.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate a task that takes time to complete
func doTask2(taskName string, duration time.Duration, wg *sync.WaitGroup) { // The function takes a task name, duration, and a pointer to a wait group
	fmt.Printf("Starting task: %s\n", taskName) // Print the name of the task being started
	time.Sleep(duration) // Simulate a delay to represent the time taken by the task
	fmt.Printf("Completed task: %s\n", taskName) // Print the name of the task that has been completed
	defer wg.Done() // Mark the task as done in the wait group -> defer ensures that wg.Done() is called when the function exits, regardless of how it exits (normal return or panic).

}

func main() {
	var wg sync.WaitGroup // Create a wait group to manage goroutines

	// Start a goroutine to perform a task
	wg.Add(1) // Add a task to the wait group
	go doTask2("Task 1", 2*time.Second, &wg) // This will run concurrently with the main function -> passing the address of the wait group to the goroutine allows it to signal when the task is done

	// Start another goroutine to perform a different task
	wg.Add(1) // Add another task to the wait group
	go doTask2("Task 2", 3*time.Second, &wg) // This will also run concurrently with the main function -> passing the aaddress of the wait group to the goroutine allows it to signal when the task is done

	// Wait for all tasks to complete
	wg.Wait() // Block until all tasks in the wait group are done
	fmt.Println("All tasks completed") // Print a message indicating that all tasks are done
}
// Output:
// Starting task: Task 1
// Starting task: Task 2
// Completed task: Task 1
// Completed task: Task 2
// All tasks completed