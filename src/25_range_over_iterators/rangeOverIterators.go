// Go added support for iterators which lets us range over pretty much anything
// TODO: read - https://go.dev/blog/range-functions

package main

import (
	"fmt"
	"iter"
	"slices"
)

// generic element type with val and next fields
type element[T any] struct {
	next *element[T]
	val  T
}

// singly linked list
type List[T any] struct {
	head, tail *element[T]
}

// to push a value to the list
func (lst *List[T]) Push(v T) { // note that we need to mention the type parameters in the method
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// returns an iterator
// iterator function takes another function as a parameter, called yield by convention(can be any arbitrary name)
func (lst *List[T]) All() iter.Seq[T] {
	// yield -> parameter
	// func (T) bool -> parameter type
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) { // call yield for every element we want to iterate over
				return
			}
		}
	}
}

func getFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		// iteration does require an underlying data structure
		// it doesn't even have to be finite
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	for e := range lst.All() {
		fmt.Println(e)
	}

	// Collect() take any iterator and collects all its values into a slice
	all := slices.Collect(lst.All())
	fmt.Println("all elements: ", all)

	for n := range getFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

// output:
// 10
// 20
// 30

// all elements:  [10 20 30]

// 1
// 1
// 2
// 3
// 5
// 8
