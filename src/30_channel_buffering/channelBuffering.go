// By default, channels are unbuffered
// i.e. they wil only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value
// Buffered channels - accept a limited number of values without a corresponding receiver for those values

package main

import "fmt"

func main() {
    messages := make(chan string, 2) // the second parameter is for buffer values

    // we can send values without being received by any receiver
    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<- messages)
    fmt.Println(<- messages)

}

// output:
// buffered
// channel