// Primary mechanism for managing state in Go is over channels
// This is a simple example of atomic counters accessed by multiple goroutines.
// https://pkg.go.dev/sync/atomic
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Uint64 // We'll use an unsigned integer to represent our (always-positive) counter.

	var wg sync.WaitGroup // WaitGroup is used to wait for the program to finish goroutines

	// To simulate concurrent updates, we'll start 50 goroutines that each increment the counter 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Even though no goroutines are writing to 'ops' in this case, we use ops.Load() to read a value from the counter even while it is being updated by other goroutines.
	fmt.Println("ops:", ops.Load())

}
