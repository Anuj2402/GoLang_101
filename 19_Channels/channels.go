/*
	What are Channels in go ?
	Channels in Go are a built-in feature used for communication between goroutines. They provide a way to safely send and receive values between concurrent tasks, allowing you to synchronize execution and share data without explicit locks or shared memory.

Channels are created using the make function:
ch := make(chan int)

You can send a value to a channel using <-:
ch <- 42 // send 42 to channel

You can receive a value from a channel

value := <-ch // receive value from channel

Channels help coordinate work and data flow in concurrent Go programs.


what are buffered channels ?

Buffered channels in Go are channels that have a specified capacity, allowing them to store multiple values without requiring an immediate receiver. 
You create a buffered channel by passing a second argument to the make function:

ch := make(chan int, 3) // Buffered channel with capacity 3

With buffered channels:

You can send up to the buffer’s capacity without blocking.
Sending blocks only when the buffer is full.
Receiving blocks only when the buffer is empty.
Buffered channels are useful when you want to decouple the timing between senders and receivers, allowing for more flexible communication between goroutines.




*/

package main

import (
    "fmt"
    // "math/rand"
    "time"
)

func processNum(numchan chan int){
    for value := range numchan{
        fmt.Println("Received in processNum:", value)
        time.Sleep(time.Second * 1)
    }
    
    
}
func sum(sumresult chan int, num1, num2 int ){
    result := num1 + num2
    sumresult <- result
}

// emailSender simulates sending emails by receiving messages from the emailChan channel.
// When all emails are sent, it signals completion by sending true to the done channel.
func emailSender(emailChan chan string, done chan bool){
    defer func(){
        done <- true
    }()

    for email := range emailChan{
        fmt.Println("Sneding email to ", email)
        time.Sleep(time.Second * 1)
    }

}


func main() {
    /*
	ch := make(chan int) // Create a channel of type int
    msgchan := make(chan string) // creating a channel of type string 
    numChan := make(chan int )
    */

    emailChan := make(chan string,100)
    done := make(chan bool)


    go emailSender(emailChan, done )

    for i:=0; i<5;i++{
        emailChan<-fmt.Sprintf("%dgmail.com", i)
    }

    close(emailChan) // this important to close the channel to avoid the deadlock 
    fmt.Println("Done sending emails")
    <-done 
    /*
        If you remove the line <-done from your code, the main function will not wait for the emailSender goroutine to finish processing all emails before exiting. As a result, your program may terminate before all emails are "sent" (i.e., before the goroutine completes its work), and you might not see all the expected output.

The line <-done blocks the main goroutine until the emailSender signals completion by sending true to the done channel. Without it, the main function could reach the end and exit immediately after printing "Done sending emails", potentially killing the goroutine before it finishes.
    */




    /*
    sumChan := make(chan int)

    go sum(sumChan, 5 ,4)

    res := <- sumChan

    fmt.Println(res)
    */
/*
   go processNum(numChan)
    for {
        numChan  <- rand.Intn(100)
    }
   
    
    // Start a goroutine to send a value to the channel
    go func() {
        ch <- 42 // Send 42 to the channel
        msgchan <- "ping"
    }()

    // Receive the value from the channel
    value := <-ch
     msg := <- msgchan
    fmt.Println("Msg from string type msg channel ", msg);

    fmt.Println("Received from channel:", value) // Output: Received from channel: 42
	*/
        time.Sleep(time.Second * 2) 

}
