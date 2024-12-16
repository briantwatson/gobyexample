// Select is a Go statement that lets a goroutine wait on multiple communication operations.

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// We use a select to await both these values simultaneously, printing each one as it arrives.
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// Total execution time is ~2 seconds since both sleeps execute concurrently.
	elapsed := time.Since(start)
	fmt.Println("\nExecution time: ", elapsed)

}
