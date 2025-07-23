// Go's slice package implements sorting for built-in and user-defined types

package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
    // sorting functions are generic, and work for any ordered built-in type
	slices.Sort(strs)
	fmt.Println("Sorted strings: ", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Sorted ints: ", ints)

    // the slices package to check if a slice is already in sorted order 
	s := slices.IsSorted(ints)
	fmt.Println("Ints sorted? ", s)

}


//output:
// $ go run src/46_sorting/sorting.go
// Sorted strings:  [a b c]
// Sorted ints:  [2 4 7]
// Ints sorted?  true