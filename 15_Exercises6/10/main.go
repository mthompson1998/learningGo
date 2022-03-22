package main

import (
	"fmt"
)

func main() {
	f := decrease()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func decrease() func() int {
	var x int
		return func() int {
		x--
		return x
	}
}