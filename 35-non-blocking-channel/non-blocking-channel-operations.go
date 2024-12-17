package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive.
	// If a value is available on messages then select will take the <-messages case with that value.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here msg cannot be sent to the messages channel, because the channel has no buffer
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// buffered channels
	messages2 := make(chan string, 1)
	signals2 := make(chan bool, 1)

	msg2 := "hi"
	signals2 <- true

	select {
	case messages2 <- msg2:
		fmt.Println("sent message", msg2)
	default:
		fmt.Println("no message sent")
	}

	// Here because both channels have a message, go randomly selects one of the two cases to execute.
	select {
	case msg := <-messages2:
		fmt.Println("received message", msg)
	case sig := <-signals2:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// Handling multiple channels with a for loop
	messages3 := make(chan string, 1)
	signals3 := make(chan bool, 1)

	messages3 <- "hi"
	signals3 <- true

	for i := 0; i < 2; i++ {
		select {
		case msg := <-messages3:
			fmt.Println("received message", msg)
		case sig := <-signals3:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
		}
	}
}
