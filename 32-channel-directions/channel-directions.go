// When using channels as function parameters, you can specify if channel is meant to only send or receive values
// This increases the type-safety of the program

package main

import "fmt"

// Only accepts channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts channel for sending and receiving values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 2)
	ping(pings, "Hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
