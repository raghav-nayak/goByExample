// Defer is used to ensure that a function call is performed later in a program's execution, usually for cleanup purposes.
// `defer` is often used where e.g. `ensure` and `finally` would be used in other languages

package main

import (
	"fmt"
	"os"
)

// program creates a file, writes to it, and then close it when we're done
func main() {
    fName := createFile("/tmp/defer.txt")
    // immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
    // This will be executed at the end of the enclosing function (main), after writeFIle has finished.
    defer closeFile(fName)
    writeFile(fName)
}

func createFile(fName string) *os.File{
    fmt.Println("creating")
    f, err := os.Create(fName)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(fName *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(fName, "data")
}

func closeFile(fName *os.File) {
    fmt.Println("closing")
    // it's important to check for errors when closing a file, even in a deferred function
    if err := fName.Close(); err != nil {
        panic(err)
    }
}

// output:
// $ go run src/49_defer/defer.go
// creating
// writing
// closing -> confirmation of the file is closed after being written
 