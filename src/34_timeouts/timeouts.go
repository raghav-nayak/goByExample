// Timeouts connect external resources or that otherwise need to bound execution time.
// Using channels and select, we can easily and elegantly implement timeouts in Go
// Here we are using time.After to achieve this

package main

import (
	"fmt"
	"time"
)

func main() {
    // buffered channel so, send in the goroutine is non-blocking
    // A common pattern to prevent goroutine leaks in case the channel is never read
    c1 := make(chan string, 1) 
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1: // awaits for the result
        fmt.Println(res)
    // time.After waits for the duration to elapse and then sends the current time on the returned channel.
    case <-time.After(1 * time.Second): // time.After awaits a value to be sent after the timeout of 1s
        fmt.Println("timeout 1") // this will be printed as time.After(1s) is received first
    }

    c2 := make(chan string, 1)
    go func ()  {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res) // this is printed as goroutine Sleep(2s) is lesser than time.After(3s) 
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}

// output:
// timeout 1
// result 2