package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // workers receive work on jobs channel
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second) // if we don't have this, all the jobs are handled by a single worker
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2 // send the results to the results channel
	}
}

func main() {
	const numsJobs = 5
	jobs := make(chan int, numsJobs)
	results := make(chan int, numsJobs)

	// to run concurrent workers; i.e. 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// sending 5 jobs
	for j := 1; j <= numsJobs; j++ {
		fmt.Println("Pushed ", j, " to jobs channel")
		jobs <- j
	}
	close(jobs) // close the channel to indicate the work we have

	// this is important to collect all the results of the work to ensure that the worker goroutines have finished
	// alternative to this is to use WaitGroup
	for a := 1; a <= numsJobs; a++ {
		<-results
	}
}

// output:
// note: even though we have 5 jobs, the concurrent workers have finished the jobs in 2s
// Pushed  1  to jobs channel
// Pushed  2  to jobs channel
// Pushed  3  to jobs channel
// Pushed  4  to jobs channel
// Pushed  5  to jobs channel
// worker 3 started job 2
// worker 2 started job 3
// worker 1 started job 1
// worker 1 finished job 1
// worker 1 started job 4
// worker 3 finished job 2
// worker 3 started job 5
// worker 2 finished job 3
// worker 1 finished job 4
// worker 3 finished job 5
