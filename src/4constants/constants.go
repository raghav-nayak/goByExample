package main

import (
	"fmt"
	"math"
)

const s string = "constant";

func main() {
	fmt.Println(s);
	const n = 500000000;

	const d = 3e20 / n; // Constant expressions perform arithmetic with arbitrary precision.
	fmt.Println(d)

	fmt.Println(int64(d)); // A numeric constant has no type until it’s given one, such as by an explicit conversion.

	fmt.Println(math.Sin(n)); // math.Sin expected float64, and explicit conversion is not required here 
	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call.
}

// output
// constant
// 6e+09
// 6000000000
// -0.28470407323754404