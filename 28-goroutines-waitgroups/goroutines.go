// Modifying goroutines to use WaitGroups
package main

import (
	"fmt"
	"sync"
)

// rather than waiting for each goroutine to finish, we use a WaitGroup to wait for all goroutines to finish.
// Pass in a WaitGroup here because the purpose of this function is to demonstrate how to run a function concurrently using goroutines.
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var wg sync.WaitGroup

	// How to call a function synchronously, usual way
	f("direct")

	// How to call a function asynchronously, using goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		f("goroutine")
	}()

	// You can also start a goroutine for an anonymous function call
	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("going")

	wg.Wait()
	fmt.Println("done")
}
