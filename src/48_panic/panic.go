// means something went  expectedly wrong.
// Mostly we use panic to fail fast on error that shouldn't occur during normal operation
// or that we aren't prepared to handle gracefully.

package main

import "os"

func main() {
    // panic("a problem")

    // A common use of panic is to abort if a function returns an error value that we dont know how to(or want to) handle.
    // e.g. the code will panic if we get an unexpected error when creating a new file.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}

// Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status
// When first panic in main fires, the program exists without reaching the rest of the code

// output 1:
// $ go run src/48_panic/panic.go
// panic: a problem

// goroutine 1 [running]:
// main.main()
//         ~/Developer/github.com/goByExample/src/48_panic/panic.go:10 +0x25
// exit status 2
// FAIL

// -------------------
// output 2: commenting first panic
// $ go run src/48_panic/panic.go

// $ ls -l /tmp/file
// -rw-rw-r-- 1 raghav raghav 0 Jul 24 22:57 /tmp/file