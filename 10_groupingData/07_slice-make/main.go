package main

import (
	"fmt"
)

func main() {
	// make(type, length, capacity)
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	//change first and last value in slice

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3223)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// add 3223 to the end
}
