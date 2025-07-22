// for more complex state, we can use mutex
// to safely access data across multiple goroutines
// we use mutex for explicit locking to synchronize access to shared state across multiple goroutines

package main

import (
	"fmt"
	"sync"
)

// container holds a map of counters
// since we want to update it concurrently from multiple goroutines, we add a mutex to synchronize access
// Mutexes must not be copied, so if this struct is passed around, it should be done by pointer
type Container struct {
    mutex sync.Mutex
    counters map[string]int
}

func (c *Container) increment(name string){
    // lock the mutex before accessing counters map
    c.mutex.Lock()
    // unlock it at the end of the function using a defer statement
    defer c.mutex.Unlock()
    c.counters[name]++
}

func main() {
    // the zero value of a mutex is usable as-is, so no initialization is needed here
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    // this function increments a names counter in a loop
    doIncrement := func (name string, n int) {
        for range n {
            c.increment(name)
        }
        wg.Done()
    }

    // to run several goroutines concurrently
    // they all access the same Container, and two of them access the same counter
    wg.Add(3)
    go doIncrement("a", 1000)
    go doIncrement("a", 1000)
    go doIncrement("b", 1000)

    // wait for the goroutines to finish
    wg.Wait()
    fmt.Println(c.counters)
}

// output:
// map[a:2000 b:1000]