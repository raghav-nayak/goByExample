package main
import "fmt"

// Go allows us to pass references to values and records within the program

func zeroVal(ival int) { // parameter is passed by value; so it does not change the value of ival
    ival = 0
}

func zeroPtr(iptr *int) { // parameter is passed by reference
    *iptr = 0
    // * here is to dereference the pointer from the its memory address to the current value at that address;
    // assigning a value to a dereferenced point changes the value at the referenced address
}

func main() {
    i := 1
    fmt.Println("initial: ", i)
    
    zeroVal(i)
    fmt.Println("zeroVal: ", i)
    
    zeroPtr(&i) // argument is passed by reference; passing memory address of i
    fmt.Println("zeroPtr: ", i)

    fmt.Println("pointer: ", &i)
}

// output:
// initial:  1
// zeroVal:  1
// zeroPtr:  0
// pointer:  0xc00009e018