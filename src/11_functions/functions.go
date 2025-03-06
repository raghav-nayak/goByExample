package main

import "fmt"

// Go requires explicit returns
func plus(a int, b int) int {
	return a + b
}

// multiple consecutive params of the same type
func plusPlus(a, b, c  int) int {
	return a + b + c
}

func main() {
	res1 := plus(1, 2)
	fmt.Println("1+2=",res1)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("1+2+3=",res2)
}