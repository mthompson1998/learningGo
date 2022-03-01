package main

import (
	"fmt"
)

type water int

var x water

func main() {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// print format Prints the type! '\n' puts value on new line
}

/*
$ go run main.go
42
main.water
*/
