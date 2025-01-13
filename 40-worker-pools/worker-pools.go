package main

import (
	"fmt"
	"time"
)

// worker function will take two channels as input. 1. jobs channel to get work and 2. results channel to send results.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2

	}
}

func main() {
	start := time.Now()

	const numJobs = 5
	// in order to use pool of workers, we need to send work and collect results.  Here we are using two channels to do that.
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// Starts 3 workers. Initially blocked as there is no work to do.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// sending 5 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// Close channel to indicate that's all the work we have.
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// this uses the synchronization feature of channels to block until all the results are available.
		<-results
	}

	// Takes about 2 seconds to finish all the jobs because each job takes 1 second to finish and we have 3 workers.
	fmt.Println("\nExecution time: ", time.Since(start))
}
