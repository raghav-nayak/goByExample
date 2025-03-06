package main

import "fmt"
import "slices"

func main() {
	var s1 []string // typed only, equals to nil, has length 0
	fmt.Println("uninitialized string slice, s1:", s1, s1 == nil, len(s1) == 0)

	s2 := make([]string, 3) // here, new slice capacity is equal to its length; but it is empty but not nil
	fmt.Println("created using make, string slice, s2:", s2, s2 == nil, len(s2), cap(s2)) 
	// len(slice) → Number of elements in the new slice.
	// cap(slice) → Number of elements from the starting index in the original slice to the end.
	// this will result same in the current case
	// the result differs when we slice the elements and assign it to another
	
	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"
	fmt.Println("after assigning values, string slice, s2:", s2, len(s2), cap(s2))
	fmt.Println("s2[2]: ", s2[2])

	// s2[3] = "d" // panic: runtime error: index out of range [3] with length 3
	
	// append returns a slice containing one or more new values and we need to accept/store in same or another slice
	s2 = append(s2, "d")
	s2 = append(s2, "e", "f")
	fmt.Println("after appending values, string slice, s2:", s2, len(s2), cap(s2))

	// copy slice
	s3 := make([]string, len(s2))
	copy(s3, s2)
	fmt.Println("after copying one slice into another, string slice, s3:", s3, len(s3), cap(s3))

	// slice operator
	s4 := s2[2:5]
	fmt.Println("sliced, string slice, s4:", s4, len(s4), cap(s4))

	s5 := s2[:5]
	fmt.Println("sliced, string slice, s5:", s5, len(s5), cap(s5))

	s6 := s2[2:]
	fmt.Println("sliced, string slice, s6:", s6, len(s6), cap(s6))

	s7 := []string{"a", "b", "c"}
	fmt.Println("declared and initialized, s7:", s7, len(s7), cap(s7))

	s8 := []string{"a", "b", "c"}
	if slices.Equal(s7, s8) {
		fmt.Println("s7 and s8 are equal")
	}

	twoDSlice := make([][]int,3) // 3 is outslice len i.e. number of rows
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("twoDSlice: ", twoDSlice)
}

// output:
// uninitialized string slice, s1: [] true true
// created using make, string slice, s2: [  ] false 3 3
// after assigning values, string slice, s2: [a b c] 3 3
// s2[2]:  c
// after appending values, string slice, s2: [a b c d e f] 6 6
// after copying one slice into another, string slice, s3: [a b c d e f] 6 6
// sliced, string slice, s4: [c d e] 3 4
// sliced, string slice, s5: [a b c d e] 5 6
// sliced, string slice, s6: [c d e f] 4 4
// declared and initialized, s7: [a b c] 3 3
// s7 and s8 are equal
// twoDSlice:  [[0] [1 2] [2 3 4]]
