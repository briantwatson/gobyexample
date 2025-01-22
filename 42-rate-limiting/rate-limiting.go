package main

import (
	"fmt"
	"time"
)

func main() {

	// basic rate limiting
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel will receive a value every 200 milliseconds.
	// This is the regulator in the rate limiting mechanism.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// By blocking on a receive from the limiter channel before serving each request,
		// we limit ourselves to 1 request every 200 milliseconds.
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Allowing short bursts of requests in rate limiting, we accomplish this by buffering our limiter channel
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we'll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range limiter { // building on the previous example
			burstyLimiter <- t
		}
	}()

	// Now we simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
