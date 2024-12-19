// Go's built-in timer and ticker features make executing code at some point in the future or repeatedly at regular intervals easy.
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Timers represent a single event in the future.
	// You tell the timer how long to wait, then it provides a channel that will be notified at that time.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer's channel C until it sends a value indicating that the timer expired.
	<-timer1.C
	fmt.Printf("Timer 1 finished after %v seconds\n", time.Since(start))

	// Adding another timer here to confirm that the timer1 is blocking the execution.
	timer3 := time.NewTimer(1 * time.Second)
	<-timer3.C
	fmt.Printf("Timer 3 finished after %v seconds\n", time.Since(start))

	// If you just wanted to wait, you could have used time.Sleep.
	// Using a timer allows you to cancel the timer before it expires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Printf("Timer 2 finished after %v seconds\n", time.Since(start))
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

	fmt.Println("\nExecution time: ", time.Since(start))
}
