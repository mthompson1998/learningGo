package main

import (
	"fmt"
)

// variables outside of function should be delcared with VAR.
// inside a function, use ':='

var y = 50

func main() {
	x := 25
	fmt.Println(x)
}

// best practice: keep scope as narrow as possible
