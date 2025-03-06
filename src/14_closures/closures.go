package main

import "fmt"

// Go supports anonymous functions which can form "closures"
// anonymous functions are inline functions without a name

func intSeq() func() int { // returns another function
    i:=0
    return func() int { // anonymous function
        i++ // this value persists between invocation of the function
        return i
    }
}

func main() {
    nextInt := intSeq() // this returns a function

    fmt.Println(nextInt())
    fmt.Println(nextInt()) // each of this invocation uses the i value from previous invocation
    fmt.Println(nextInt())
    
    newNextInt := intSeq()
    fmt.Println(newNextInt())

} 

// output:
// 1
// 2
// 3
// 1