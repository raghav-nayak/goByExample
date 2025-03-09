// we use use for and range to iterate over values received from a channel

package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue) // close the channel

    // iterates over each element as it's received from the channel
    // iteration terminates after receiving 2 elements as we closed the channel
    // note: it is possible to close a non-empty channel but still have the remaining values be received
    for elem := range queue {
        fmt.Println(elem)
    }
}

// output:
// one
// two