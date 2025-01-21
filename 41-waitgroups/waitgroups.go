package main

import (
	"fmt"
	"sync"
	"time"
)

// function we'll run every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// simulating an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Used to wait for all goroutines to finish. If a WaitGroup is explicitly passed to functions, it should be done as a pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Add 1 to the WaitGroup for each goroutine
		wg.Add(1)
		go func() {

			// Decrement the WaitGroup counter when the goroutine completes. This should be done in a defer statement to ensure it is always called.
			// This way the worker function can be called without needing to know about the WaitGroup.
			defer wg.Done()
			worker(i)
		}()
	}

	// block until waitgroup counter is 0
	wg.Wait()
}
