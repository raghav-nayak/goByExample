// when using channels as function parameters, we can specify if a channel is meant to only send or receive values
// this specificity increases the type-safety of the program

// chan<-  sends
// <-chan  receive

package main

import "fmt"

// this function only accepts a channel for sending values
// it would be a compile-time error to try to receive on this channel
func ping(pings chan<- string, msg string) {
    pings <- msg
}


// this function takes two channels as parameters
// pings for receiving
// pongs for sending
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}

// output:
// passed message