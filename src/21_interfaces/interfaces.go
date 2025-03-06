// interfaces are named collections of method signatures
// read: https://research.swtch.com/interfaces

package main

import (
	"fmt"
	"math"
)

// interface definition
type geometry interface {
    area() float64
    perimeter() float64
}

// implementation of interface by implementing all the methods in the interfaces
type rectangle struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rectangle) area() float64 {
    return r.width * r.height
}

func (r rectangle) perimeter() float64 {
    return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) { // the parameter is an interface type
    fmt.Println(g)
    fmt.Println(g.area()) // we can call the methods in the named interface
    fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
    // to get the runtime type of the interface value
    // below is the type assertion; check 7_switch
    if c, ok := g.(circle); ok { 
        fmt.Println("circle with radius", c.radius)
    }
}

func main() {
    r := rectangle{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

    detectCircle(r)
    detectCircle(c)
}

// output:
// {3 4}
// 12
// 14

// {5}
// 78.53981633974483
// 31.41592653589793

// circle with radius 5