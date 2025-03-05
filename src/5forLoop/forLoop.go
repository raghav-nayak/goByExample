package main

import "fmt"

func main() {
	i := 1;

	for i <= 3 {
		fmt.Println("basic type:", i)
		i = i+1;
	}

	for j := 0; j < 3; j++ {
		fmt.Println("classic:", j);
	}

	for i := range 3 {
		fmt.Println("range:", i)
	}

	for {
		fmt.Println("without condition");
		break;
	}

	for n := range 6 {
		if n%2 == 0 {
			continue;
		}
		fmt.Println("range with continue:", n)
	}
}

// output:
// basic type: 1
// basic type: 2
// basic type: 3
// classic: 0
// classic: 1
// classic: 2
// range: 0
// range: 1
// range: 2
// without condition
// range with continue: 1
// range with continue: 3
// range with continue: 5