package main

import (
	"fmt"
	"time"
)

func main() {

	// Suppose we're executing an external call that returns its result on a channel c1 after 2 seconds.
	// Note that the channel is buffered, so the send in the goroutine is non-blocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Here's the select implementing a timeout
	// res:= <-c1 awaits result
	select {
	case res := <-c1:
		fmt.Println(res)
	// <-time.After awaits a value to be sent after the timeout of 1 second.
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
