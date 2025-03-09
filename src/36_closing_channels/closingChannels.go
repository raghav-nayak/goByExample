// Closing a channel indicates that no more values will be sent on it.
// useful to communicate completion to the channel's receivers

package main

import "fmt"

func main() {
	jobs := make(chan int, 5) // buffered channel
	done := make(chan bool)

    // jobs channel -> used to communicate work to be done from the main() goroutine to a worker goroutine
    // done channel -> used when we have no more jobs for the worker we close the jobs channel

	go func() {
		for { // to run infinitely
			j, more := <-jobs // 2-value of form of receive; 
            // more will false if jobs is closed and all the values in the channel have already been received
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // to notify on done when we've worked all our jobs; this value is not important here
				return
			}
		}
	}()

    // send three jobs to the worker over the jobs channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // close the channel
	fmt.Println("sent all jobs")
	
    // await the worker using the synchronous approach
    // block until we receive a notification from the worker on the channel
    // see the output: with and without this line
    // the value is not important here; 
    <-done 

    // reading from the closed channel succeeds immediately, returning zero value of the underlying type
    // the second return value:
    //      true - the value received was delivered by a successful send operation to the channel
    //      false - the value is a zero value generated because the channel is closed and empty
    _, ok := <-jobs 
	fmt.Println("received more jobs:", ok)
}

// output:

// with <-done:
// sent job 1
// sent job 2
// sent job 3
// sent all jobs
// received job 1
// received job 2
// received job 3
// received all jobs
// received more jobs: false

// without <-done:
// sent job 1
// sent job 2
// sent job 3
// sent all jobs
// received more jobs: true
