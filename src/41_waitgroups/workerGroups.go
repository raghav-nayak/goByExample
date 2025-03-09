// WaitGroup is used to wait for multiple goroutines to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// to wait for all the goroutines launched here to finish
	// if a WaitGroup is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the WaitGroup counter for each goroutine

		// wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done
		// this will make sure that the worker itself does not have to be aware of the concurrency primitives involved in its execution
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// block until the WaitGroup counter goes back 0; all the workers notified they'are done
	wg.Wait() // this is important
	// note: this approach has no straightforward way to propagate errors from workers.
	// In advance cases, use errgroup package
}

// output:
// Worker 5 starting
// Worker 1 starting
// Worker 2 starting
// Worker 3 starting
// Worker 4 starting
// Worker 4 done
// Worker 2 done
// Worker 3 done
// Worker 5 done
// Worker 1 done
