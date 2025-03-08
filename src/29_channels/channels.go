// channels - the pipes that connect concurrent goroutines
// we can send values from one goroutine and receive those values into another goroutines

package main

import "fmt"

func main() {
    // create channel - make(chan val-type)
    // channels are typed by the values they convey
    messages := make(chan string)

    go func() {
        messages <- "ping" // send a value to a channel
    }()
    message := <- messages // receive a value from a channel
    // By default, send and receive block until both the sender and receiver are ready
    // this property allows us to wait at the end of the program for the message without 
    // having to use any other synchronization

    fmt.Println(message)
}

// output:
// ping