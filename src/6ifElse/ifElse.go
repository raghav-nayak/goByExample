package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}
	
	// num is declared and available in the current and subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 { // branches are not needed but parentheses
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// note: there is no ternary if in Go
}

// output:
// 7 is odd
// 8 is divisible by 4
// either 8 or 7 are even
// 9 has 1 digit