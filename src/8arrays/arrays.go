package main

import "fmt"

func main() {
	var a [5]int // By default for int type, it is zero-valued
	fmt.Println("declare a:", a)
	a[4] = 100
	fmt.Println("after setting, a:", a)
	fmt.Println("get a:", a[4])
	fmt.Println("len of a:", len(a))

	b := [5]int{1,2,3,4,5} // explicitly specifying the number of elements
	fmt.Println("declare and initialize explicitly, b:", b)

	c := [...]int{1,2,3,4,5} // the compiler count the number of elements
	fmt.Println("declare and initialize implicitly, c:", c)

	d := [...]int{100, 3:400, 500}
	fmt.Println("declare and initialize skipping between elements, d:", d)

	var twoDArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDArray[i][j] = i + j
		}
	}
	fmt.Println("twoDArray: ", twoDArray)

	twoDArrayInit := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("twoDArrayInit: ", twoDArrayInit)
}

// output:
// declare a: [0 0 0 0 0]
// after setting, a: [0 0 0 0 100]
// get a: 100
// len of a: 5
// declare and initialize explicitly, b: [1 2 3 4 5]
// declare and initialize implicitly, c: [1 2 3 4 5]
// declare and initialize skipping between elements, d: [100 0 0 400 500]
// twoDArray:  [[0 1 2] [1 2 3]]
// twoDArrayInit:  [[1 2 3] [4 5 6]]