// Go provides an explicit, separate return value for errors.
// This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.
// Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.
// TODO read - https://pkg.go.dev/errors

package main

import (
	"errors"
	"fmt"
)

// by convention, error is the last return value and have type error(built-in interface)
func f(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("cannot work with 42") // constructs a basic error value with the given error message
    }
    return arg + 3, nil // nil to indicate no error
}

// sentinel error - predeclared variable used to signify a specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("cannot boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {
        return fmt.Errorf("making tea: %w", ErrPower) // wrap errors which higher-level errors; can be queried using errors.Is and errors.As
    }
    return nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f(i); e != nil { // inline error check in if 
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    // errors.Is reports whether any error in err's tree matches target.
    // check that a given error(or any error in tis chain) matches a specific error value
    // useful with wrapped or nested errors, allows to identify specific error types or sentinel errors in a chain of errors
    for i := range 5 { 
        if err := makeTea(i); err != nil {
            if errors.Is(err, ErrOutOfTea) { 
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
        }
        fmt.Println("Tea is ready for ", i)
    }
}

// output:
// f worked: 10
// f failed: cannot work with 42

// Tea is ready! for  0
// Tea is ready! for  1
// We should buy new tea!
// Tea is ready! for  2
// Tea is ready! for  3
// Now it is dark
// Tea is ready! for  4