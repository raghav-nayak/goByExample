package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println("a:",a, " b:", b)

	_, c := vals()
	fmt.Println("c:",c)
}

// output:
// a: 3  b: 7
// c: 7
