// Go structs are typed collections of fields
// useful for grouping data together to form records

package main

import "fmt"

type person struct {
    name string
    age int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 18
    return &p 
    // we can safely return a pointer to a local variable
    // Go language is a garbage collected(GC) language
    // it will only be cleaned up by the GC when there are no active references to it
}

func main() {
    fmt.Println(person{"Raja", 20}) // to create a new struct
    fmt.Println(person{name:"Rank", age:24}) // we can name the fields while initializing
    fmt.Println(person{name:"Jai"}) // omitted values are zero-valued
    fmt.Println(&person{name:"Veeru", age:25}) // yields a pointer to the struct
    fmt.Println(newPerson("Basanti")) // to encapsulate new struct creation in constructor functions

    s := person{name:"Gabbar", age:45}
    fmt.Println("name:", s.name) // to access struct fields

    sp := &s
    fmt.Println("age using pointer:", sp.age)
    
    sp.age = 40 // structs are mutable
    fmt.Println("age after updating:", sp.age)

    // anonymous struct type; commonly used in table-driven tests
    dog := struct {
        name string
        isGood bool
    }{
        "Melon",
        true,
    }
    fmt.Println("dog:", dog)
}

// output:
// {Raja 20}
// {Rank 24}
// {Jai 0}
// &{Veeru 25}
// &{Basanti 18}
// name: Gabbar
// age using pointer: 45
// age after updating: 40
// dog: {Melon true}
