// We can use channels to synchronize execution across goroutines.

package main

import (
	"fmt"
	"time"
)

// example of using a blocking receive to wait for a goroutine to finish
// we can use WaitGroup for waiting for multiple goroutines to finish
func worker(done chan bool) { // done channel is used to notify another goroutine that this function's work is done
	fmt.Println("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true // send a value to notify that we're done 
}

func main() {
    done := make(chan bool, 1)

    go worker(done) // start a goroutine, giving it to the channel to notify on

    <- done // block until we receive a notification from the worker on the channel
    // if we remove this line, the program would exit before the worker even started
}

// output:
// working...
// done
