package main

import (
	"fmt"
)

var y = 42

// Delcare the variable with the IDENTIFIER 'z'
// is of TYPE string
var z string = "shaken, not stirred"
var a string = `james says: "shaken, not stirred"`

//raw explicit, grabs all info between ` ` backticks

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
