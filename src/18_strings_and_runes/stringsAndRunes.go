// Go string is a read-only slice of bytes.
// Go standard lib treats strings specially - as containers of text encoded in UTF-8
// In other languages, strings are made of chars
// In Go, the concept of character is called a rune - integer that represents a Unicode code point

package main

import (
	"fmt"
	"unicode/utf8"
)

// import "unicode/utf8"

func main() {
    const s = "ನಮಸ್ಕಾರ" // Go string literals are UTF-8 encoded text
    fmt.Println("s: ",s, " len: ",len(s)) // len returns the length of raw bytes stored within
    for i:=0; i < len(s); i++ { // looping over raw bytes
        fmt.Printf("%x ", s[i]) // prints hex value of all bytes
    }
    fmt.Println()

    // decoded each rune sequentially at the run-time
    // the len of string and this count are not same
    fmt.Println("Rune count: ", utf8.RuneCountInString(s)) 
    
    const s2= "raghav"
    fmt.Println("Rune count: ", utf8.RuneCountInString(s2)) // this is same as len of string

    for idx, runeVal := range s {
        fmt.Printf("%#U start at %d\n", runeVal, idx) // to decode each rune
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeVal, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeVal, i)
        w = width

        examineRune(runeVal)
    }
}

func examineRune(r rune) { // rune as parameter
    if r == 'ನ' || r == 'ಕ' || r == 'ರ' {
        fmt.Println("found ನ or ಕ or ರ")
    } else {
        fmt.Println("found other runes")
    }
}

// output:
// s:  ನಮಸ್ಕಾರ  len:  21
// e0 b2 a8 e0 b2 ae e0 b2 b8 e0 b3 8d e0 b2 95 e0 b2 be e0 b2 b0 
// Rune count:  7
// Rune count:  6
// U+0CA8 'ನ' start at 0
// U+0CAE 'ಮ' start at 3
// U+0CB8 'ಸ' start at 6
// U+0CCD '್' start at 9
// U+0C95 'ಕ' start at 12
// U+0CBE 'ಾ' start at 15
// U+0CB0 'ರ' start at 18

// Using DecodeRuneInString
// U+0CA8 'ನ' starts at 0
// found ನ or ಕ or ರ
// U+0CAE 'ಮ' starts at 3
// found other runes
// U+0CB8 'ಸ' starts at 6
// found other runes
// U+0CCD '್' starts at 9
// found other runes
// U+0C95 'ಕ' starts at 12
// found ನ or ಕ or ರ
// U+0CBE 'ಾ' starts at 15
// found other runes
// U+0CB0 'ರ' starts at 18
// found ನ or ಕ or ರ