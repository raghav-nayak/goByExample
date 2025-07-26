// The std library's strings package provides many useful string-related functions.
// These are  functions from the package and not methods on the string object itself.
// So, we need to pass the string object as the first argument to the function.

package main

import (
	"fmt"
	"strings"
)

// creating alias
var print = fmt.Println

func main() {
	print("Contains: ", strings.Contains("test", "es"))
	print("Count: ", strings.Count("test", "t"))
	print("HasPrefix: ", strings.HasPrefix("test", "te"))
	print("HasSuffix: ", strings.HasSuffix("test", "st"))
	print("Index: ", strings.Index("test", "e"))
	print("Join: ", strings.Join([]string{"a", "b"}, "-"))
	print("Repeat: ", strings.Repeat("a", 5))
	print("Replace: ", strings.Replace("foo", "o", "0", -1))
	print("Replace: ", strings.Replace("foo", "o", "0", 1))
	print("Split: ", strings.Split("a-b-c-d-e", "-"))
	print("ToLower: ", strings.ToLower("TEST"))
	print("ToUpper: ", strings.ToUpper("test"))
}

// output:
// $ go run src/51_string_functions/string_functions.go
// Contains:  true
// Count:  2
// HasPrefix:  true
// HasSuffix:  true
// Index:  1
// Join:  a-b
// Repeat:  aaaaa
// Replace:  f00
// Replace:  f0o
// Split:  [a b c d e]
// ToLower:  test
// ToUpper:  TEST
