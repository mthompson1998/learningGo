package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := func() int{
		return 451
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x())
}

func foo() string {
	return "hello world"
}

//func bar() func() int {
//	return func() int {
//		return 451
//	}()
