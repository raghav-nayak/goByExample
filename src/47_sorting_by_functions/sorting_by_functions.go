// To sort a collection by something other than its natural order, e.g. to sort strings by their length instead of
// alphabetically, we can use custom sorts in Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

    // implementation of a comparison function for string lengths using cmp.Compare 
    // refer https://pkg.go.dev/cmp#Ordered
	lengthCompare := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

    // invoking slice.SortFunc with above custom comparison function to sort fruits by name length
	slices.SortFunc(fruits, lengthCompare)
	fmt.Println("Sorted by length, fruits:", fruits)

    type Person struct {
        name string
        age int
    }

    people := []Person{
        Person{name: "Jane", age: 47},
        Person{name: "John", age: 31},
        Person{name: "Alice", age: 29},
        Person{name: "Bob", age: 35},
    }

    // to sort a slice of values that aren't built-in types
    // if the struct is large, we can use slice to contain *Person instead and adjust the sorting the function accordingly
    slices.SortFunc(people, 
        func(a, b Person) int {
            return cmp.Compare(a.age, b.age)   
        },
    )
    fmt.Println("Sorted by age, people:", people)
}


// output:
// $ go run src/47_sorting_by_functions/sorting_by_functions.go
// Sorted by length, fruits: [kiwi peach banana]
// Sorted by age, people: [{Alice 29} {John 31} {Bob 35} {Jane 47}]