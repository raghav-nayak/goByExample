// Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution instead.

// usecase:
// A server wouldn't want to crash if one of the client connections exhibits a critical error.
// Instead, the server would want to close that connection and continue serving other clients.
// If fact, this is what Go's net/http does by default for HTTP servers.

package main

import "fmt"

// this function panics
func mayPanic() {
    panic("a problem")
}

func main() {

    // recover must be called within a deferred function. When the enclosing function panics, the defer will activate
    // and a recover call within it will catch the panic
    defer func() {
        // the return value of recover is the error raised in the call to panic
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()")
}

// Output:
// $ go run src/50_recover/recover.go
// Recovered. Error:
//  a problem