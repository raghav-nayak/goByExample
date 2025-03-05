package main

import "fmt"

func main() {
	fmt.Println("go" + "lang"); // string concatenation
	fmt.Println("1+1 =", 1+1); // integer addition
	fmt.Println("7.0/3.0 =", 7.0/3.0); // float division
	fmt.Println(true && false); // logical AND
	fmt.Println(true || false); // logical OR
	fmt.Println(!true); // logical NOT
}

// output:
// golang
// 1+1 = 2
// 7.0/3.0 = 2.3333333333333335
// false
// true
// false 