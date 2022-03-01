package main

import (
	"fmt"
)

func main() {
	var x [5] int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
}

/*
go run main.go
[0 0 0 0 0]
[0 0 0 42 0]
*/

//don't really want to use arrays.... use SLICES!
//ease of programming, fast compilation.