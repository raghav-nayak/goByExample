package main

import "fmt"

type point struct {
	x, y int
}

func main() {
    p := point{1,2}

    // prints an instance of the struct
    fmt.Printf("struct1: %v\n", p) // struct1: {1 2}

    // prints struct's field names
    fmt.Printf("struct2: %+v\n", p) // struct2: {x:1 y:2}
    
    // prints Go syntax representation of the value; i.e. the source code snippet that would product that value
    fmt.Printf("struct3: %#v\n", p) // struct3: main.point{x:1, y:2}

    // the type of value
    fmt.Printf("type: %T\n", p) // type: main.point
    
    // formatting bool
    fmt.Printf("bool: %t\n", true) // bool: true

    // base-10 representation of integer
    fmt.Printf("int: %d\n", 123) // int: 123

    // binary representation of integer
    fmt.Printf("bin: %b\n", 14) // bin: 1110

    // ASCII character corresponds to the given integer
    fmt.Printf("char: %c\n", 65) // char: A

    // hex encoding
    fmt.Printf("hex: %x\n", 456) // hex: 1c8

    // basic decimal formatting
    fmt.Printf("float1: %f\n", 78.9) // float1: 78.900000

    // decimal formatting for scientific notation
    fmt.Printf("float2: %e\n", 123400000.0) // float2: 1.234000e+08

    // decimal formatting for scientific notation
    fmt.Printf("float2: %E\n", 123400000.0) // float2: 1.234000E+08

    // basic string formatting
    fmt.Printf("str1: %s\n", "\"string\"") // str1: "string"

    // to print double-quoted string
    fmt.Printf("str2: %q\n", "\"string\"") // str2: "\"string\""

    // base-16 with two output chars per byte of input
    fmt.Printf("str3: %x\n", "hex this") // str3: 6865782074686973

    // a representation of a pointer
    fmt.Printf("pointer: %p\n", &p) // pointer: 0xc000184040

    // To specify the width of an integer, use a number after the % in the verb.
    // By default the result will be right-justified and padded with spaces.
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // width1: |    12|   345|
}
