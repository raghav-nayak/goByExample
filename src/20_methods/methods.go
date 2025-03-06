// Go supports methods defined on struct types
// methods can be defined for either pointer or value receiver types
// Go automatically handles the conversion between values and pointer for method calls

package main

import "fmt"

type rectangle struct {
	width, height int
}

// method has a receiver type of *rectangle
// this avoid copying on method calls or to allow the method to mutate receiving struct
func (r *rectangle) area() int {
	return r.width * r.height
}

// value type receiver
func (r rectangle) perimeter() int {
	return 2 * (r.width + r.width)
}

func main() {
	r := rectangle{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}

// output:
// area: 50
// perimeter: 40
// area: 50
// perimeter: 40
