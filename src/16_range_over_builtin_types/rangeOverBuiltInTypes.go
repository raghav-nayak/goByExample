package main

import "fmt"

func main() {
    nums := []int{2,3,4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // range on arrays and slices provides both the index and value for each entry
    for i, v := range nums {
        fmt.Println("index:", i, "value:", v)
    }

    // range on map iterates over key/value pairs.
    kvs := map[string]string {"a": "apple", "b": "bada apple"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n",k,v)
    }

    // range can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key:",k)
    }

    // range on strings iterates over Unicode code points
    for i,c := range "go" {
        fmt.Println(i, c)
    }
}

// output:
// sum: 9
// index: 0 value: 2
// index: 1 value: 3
// index: 2 value: 4
// a -> apple
// b -> bada apple
// key: b
// key: a
// 0 103 -> index of rune, rune
// 1 111