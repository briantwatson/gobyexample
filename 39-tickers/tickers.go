// Tickers are for when you want to do something repeatedly at regular intervals.

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)

	// This goroutine will run every 500ms
	// Tickers use a similar mechanism to timers: a channel that is sent values.
	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Tickers can be stopped like timers.
	ticker.Stop()
	done <- true
}
