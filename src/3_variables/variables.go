package main

import "fmt"

func main() {
	var a = "initial";
	fmt.Println(a);

	var b,c int = 1,2 ;
	fmt.Println(b, c)

	var d = true;
	fmt.Println(d);

	var e int; //zero-valued: variables declared without a corresponding initialization are zero-valued
	fmt.Println(e);

	f := "apple";
	fmt.Println(f);
}

// output:
// initial
// 1 2
// true
// 0
// apple