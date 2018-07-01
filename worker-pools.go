package main

import (
	"fmt"
	"time"
)

// In this example we’ll look at how to implement a worker pool using goroutines and channels.

// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and send the corresponding results on results.
// We’ll sleep a second per job to simulate an expensive task.=
func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", id, " started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker:", id, "finished job", j)
		results <- j * 2
	}
}

// WorkerPool Function to illustrate an example that create worker pool using goroutines and channels.
func WorkerPool() {

	// In order to use our pool of workers we need to send them work and collect their results.
	// We make 2 channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}

	// Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
	// Our running program shows the 5 jobs being executed by various workers.
	// The program only takes about 2 seconds despite doing about 5 seconds of total work
	// because there are 3 workers operating concurrently.

}
