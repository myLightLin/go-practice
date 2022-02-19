package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	// create tow goroutines
	go doWork("A")
	go doWork("B")

	// give them time to run
	time.Sleep(1 * time.Second)

	// shutdown the system
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	// wait for the goroutines to finish
	wg.Wait()
}

// dowork is a function that does some work
// shutdown flag is used to tell the function to exit
func doWork(name string) {
	defer wg.Done()

	// infinite loop
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// check if shutdown flag is set
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
