package main

import "SafariGo/util"
import "fmt"

// var anotherMessage string = "More text"
var anotherMessage = "More text"

const (
	//joker = 99
	hearts   int = 1 << iota
	diamonds     = 1 << iota
	clubs        = 1 << iota
	spades       = 1 << iota
)

// m2 := "xxx" // local variables only
func main() {
	var m2 string // "uninitialized" are initialized :) to a "zero" value, for string that's ""
	m1 := "Hello"
	fmt.Println("Hello Go World!")
	fmt.Println(util.Message)
	fmt.Println(anotherMessage)
	fmt.Println(m1)
	fmt.Printf("m2 is %#v and type is %T\n", m2, m2)

	// constants, e.g. "const" declarations and numeric literals
	// are handled by the compiler to "very high" range & precision
	var num int8 = 12345678901234567890 / 10000000000000000000
	fmt.Println(num)
}
