package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as " );
	// basic switch
	switch i {
	case 1:
		fmt.Println("one");
	case 2:
		fmt.Println("two");
	case 3:
		fmt.Println("three");
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // multiple expressions in a single statement
		fmt.Println("It's the weekend");
	default:
		fmt.Println("It's a weekday");
	}

	t := time.Now()
	switch { // switch without expression; alternative way for if-else
	case t.Hour() < 12:
		fmt.Println("It's before noon");
	default:
		fmt.Println("It's after noon");
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) { // comparing type instead of expression
		case bool:
			fmt.Println("I'm a bool");
		case int:
			fmt.Println("I'm an int");
		default:
			fmt.Printf("Don't know type %T\n", t);
		}
	}
	whatAmI(true);
	whatAmI(1);
	whatAmI("hey");
}

// output:
// Write 2 as two
// It's a weekday
// It's after noon
// I'm a bool
// I'm an int
// Don't know type string