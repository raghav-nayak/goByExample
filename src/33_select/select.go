// Go's select lets us wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful of Go.

package main

import (
	"fmt"
	"time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        // receives a value after some amount of time 
        // to simulate blocking RPC operations executing in concurrent goroutines
        time.Sleep(1 * time.Second)
        c1 <- "one" 
    }()

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two" 
    }()

    // total execution time is not ~2 seconds as both Sleeps execute concurrently

    for i := 0; i < 2; i++ {
        select { // use select to await both of these values simultaneously, printing each one as it arrives
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        } 
    }
}

// output:
// received one
// received two