/*
what are mutex in go ?
A mutex (short for "mutual exclusion") in Go is a synchronization primitive used to protect shared data from being accessed by multiple goroutines at the same time.
It ensures that only one goroutine can access the critical section of code at a time, preventing race conditions.

In Go, you use the sync.Mutex type:

import "sync"

var mu sync.Mutex

// To lock:
mu.Lock()
// critical section (only one goroutine at a time)
mu.Unlock()


mu.Lock() locks the mutex.
mu.Unlock() unlocks it.

Mutexes are essential for safe concurrent programming when goroutines share data.

*/

package main

import (
	"fmt"
	"sync"
)

type post struct {
	mu    sync.Mutex
	views int 
}

func (p *post) inc(wg *sync.WaitGroup){
	defer func (){
		// best practice is to keep the unlock function inside a defer function so that anyhow it will get called 
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views +=1;
	
}

func main(){
	var wg sync.WaitGroup // declaring a wait group 
	myPost := post{views: 0}


	for i:=0;i<100;i++{
		wg.Add(1)
		go myPost.inc(&wg) // output will be 100
	}

	//myPost.inc() // increment it by 1 
	wg.Wait()

	fmt.Println(myPost.views) // output -> 1

}