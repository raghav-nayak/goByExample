// goroutine -> a lightweight thread of execution

package main

import (
	"fmt"
	"time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    f("direct") // synchronous

    go f("goroutine") // goroutine will execute concurrently with the calling one

    go func(msg string) { // goroutine with an anonymous function call
        fmt.Println(msg)
    }("going")

    // both above functions are running asynchronously in separate goroutines

    // crud way is to use Sleep; 
    // robust approach: we can use WaitGroup to wait for them to finish
    time.Sleep(time.Second)
    fmt.Println("done")
}

// output:
// -- blocking call
// direct : 0
// direct : 1
// direct : 2
// -- this output can be in any order as these are running concurrently by the Go runtime
// going
// goroutine : 0
// goroutine : 1
// goroutine : 2
// done