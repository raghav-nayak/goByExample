package main

import (
	"errors"
	"fmt"
)

// custom error type usually has the suffix "Error"
type argError struct {
    arg int
    message string
}

// this method makes argError implement the error interface
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "cannot work with it"} // return custom error
    }
    return arg+3, nil
}

func main() {
    _, err := f(42)
    var ae * argError
    // errors.As is a more advanced version of errors.Is
    // checks a given error(or any error in its chain) matches a specific error type 
    // and converts to a value of that type, returning true
    // If there is no match, it returns false
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err does not match argError")
    }
}

// output:
// 42
// cannot work with it