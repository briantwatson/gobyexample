package main

import (
	"fmt"
	"time"
)

// This worker function will run concurrently in a goroutine. The done channel will be used to notify another goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify we're done
	done <- true
}

func main() {

	// Start a worker goroutine, giving it the done channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the done channel
	<-done

	// If you removed the <-done line from this program, the program would exit before the worker even started.
}
