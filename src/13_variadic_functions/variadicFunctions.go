package main

import "fmt"

// take arbitraty number of ints as args
func sum(nums ...int) {
	fmt.Println("\n", nums)

	total := 0
	// within the function, nums is equivalent to []int
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	// sum(nums) // Error: cannot use nums (variable of type []int) as int value in argument to sum
	sum(nums...) // as the function expects int numbers, we need to use ... to call the function
}

// output:
//  [1 2]
// 3

//  [1 2 3]
// 6

//  [1 2 3 4]
// 10