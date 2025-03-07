// Go added support for generics in 1.18; also known as type parameters
// read: https://go.dev/blog/deconstructing-type-parameters

package main

import "fmt"

// takes two args: a slice of any comparable type and an element of that type
// comparable constraint -> we can compare values of this type with the == and != operators
// this is same function present in the slices.Index
func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i // returns the first occurrence of v in s
		}
	}
	return -1
}

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

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"nishant", "rahul", "rajeev"}

	// when invoking generic functions, we can ofter rely on type inference.
	// while invoking , we do not have to specify any types for S or E; compiler infers automatically
	fmt.Println("index of rahul:", SliceIndex(s, "rahul"))

	// this example shows specifying the types explicitly while invoking
	_ = SliceIndex[[]string, string](s, "rahul")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println("list contents:", lst.AllElements())
}


// output:
// index of rahul: 1
// list contents: [10 20 30]