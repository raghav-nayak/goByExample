// Go's built-in timer and ticker feature help us execute the Go code at some oin the future, or repeatedly at some interval

package main

import (
	"fmt"
	"time"
)

func main() {
    // Timers represents a single event in the future
    // you can specify the timer how log you want to wait, and
    // it provides a channel that will be notified at that time
    timer1 := time.NewTimer(2 * time.Second) // waits for 2s

    // blocks on the timer's channel C util it sends a value indicating that the timer fired
    <-timer1.C
    fmt.Println("timer 1 fired")

    // we can use time.Sleep() if we want to wait
    // a timer comes handy when we want to cancel the time before it fires
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("timer 2 fired")
    }()
    stop2 := timer2.Stop() // stops the timer
    if stop2 {
        fmt.Println("timer 2 stopped")
    }

    // this is just to demonstrate whether the timer is fired or stopped
    time.Sleep(2 * time.Second)
}

// output:
// timer 1 fired
// timer 2 stopped

// note: The first timer will fire ~2s after we start the program, 
// but the second should be stopped before it has a chance to fire.
 