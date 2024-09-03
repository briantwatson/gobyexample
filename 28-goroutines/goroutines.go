// A goroutine is a lightweight thread of execution.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// How to call a function synchronously, usual way
	f("direct")

	// How to call a function asynchronously, using goroutine
	// This will execute the function concurrently with the calling function, main()
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
