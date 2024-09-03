// Channels are pipes that connect concurrent goroutines.
package main

import "fmt"

func main() {
	messages := make(chan string)

	// Send a value into a channel using the <- syntax
	// Here we send "ping" to the messages channel we made above using an anonymous goroutine
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
