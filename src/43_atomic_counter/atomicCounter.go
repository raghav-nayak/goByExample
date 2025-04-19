// The primary mechanism for managing state in Go is communication over channels
// There are other options for managing state
// sync/atomic package for atomic counter accessed by multiple goroutines

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1) // Add atomically adds delta to x and returns the new value.
			}
            wg.Done()
		}()
	}

    // Wait blocks until the [WaitGroup] counter is zero.
    // i.e. Wait until all the goroutines are done.
    wg.Wait() 

    // Load atomically loads and returns the value stored in x.
    // it’s safe to atomically read a value even while other goroutines are (atomically) updating it.
    fmt.Println("ops: ", ops.Load()) 

    // Note: 
    // We expect to get exactly 50,000 operations. Had we used a non-atomic integer and incremented it with ops++, 
    // we’d likely get a different number, changing between runs, because the goroutines would interfere with each other.
    // Moreover, we’d get data race failures when running with the -race flag.
}

// output:
// ops:  50000
