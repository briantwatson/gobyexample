// Closing a channel indicates that no more values will be sent on it.
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine.  It repeatedly receives from jobs with j, more := <-jobs.
	// In this special 2-value form of receive, the `more` value will be false if jobs has been closed and all values in the channel will have already been received.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	// Reading from a closed channel will return the zero value immediately.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
