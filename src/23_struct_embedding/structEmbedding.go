// Go supports embedding of structs and interfaces to the express more a more seamless composition of type
// it is different from embed go directive introduced in 1.16 to embed files and folders into the application directory

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num: %v", b.num)
}

type container struct {
	base // embed: a field without a name
	str  string
}

func main() {
	co := container{
		base: base{ // we need to initialize embedding explicitly; embed type name serves as a field name
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)                       // num is accessed using co.num
	fmt.Println("another way to print num: ", co.base.num)                      // num is accessed using full path of embedded type name
	fmt.Println("yet another way to print num, describe: ", co.base.describe()) // num is accessed using a method of embedded type

	type describer interface {
		describe() string
	}
    
    // TODO: Need further study
	var d describer = co // embedding structs with methods may be for interface implementations onto other structs
	fmt.Println("describer: ", d.describe())
}

// output:
// co={num: 1, str: some name}
// another way to print num:  1
// yet another way to print num, describe:  base with num: 1
// describer:  base with num: 1
