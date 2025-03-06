package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps - built-in associated data type
	// similar to hashes or dicts in other languages
	// map creation syntax: make(map[key-type]val-type).
	m := make(map[string]int)
	fmt.Println("after declaring, map:", m)

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println("after assigning values, map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"] // for non-existing key, default value is 0 as the value is of type int
	fmt.Println("v3:", v3)

	v4, v4Present := m["k2"] // for non-existing key, default value is 0 as the value is of type int
	fmt.Println("v4:", v4, " v4Present:",v4Present)

	fmt.Println("len:", len(m))

	delete(m, "k2") // to remove a single key-value pair
	fmt.Println("after deleting k2, map:", m)

	clear(m) // to remove all the key-value pairs
	fmt.Println("after clearing, map:", m)

	// optional second return value to check if the key-value present or not
	// Useful to distinguish between the key present and the default value(0 or "")
	// _ means ignore the value 
	_, prs := m["k2"] 
	fmt.Println("prs:", prs)

	m2 := map[string]string{"foo": "a", "bar": "b"}
	fmt.Println("after declaring and initializing, map:", m2) // while displaying, we get sorted by key

	m3 := map[string]string{"foo": "a", "bar": "b"}
	if maps.Equal(m2, m3) {
		fmt.Println("maps are equal")
	}
}

// output:
// after declaring, map: map[]
// after assigning values, map: map[k1:10 k2:20]
// v1: 10
// v3: 0
// v4: 20  v4Present: true
// len: 2
// after deleting k2, map: map[k1:10]
// after clearing, map: map[]
// prs: false
// after declaring and initializing, map: map[bar:b foo:a]
// maps are equal