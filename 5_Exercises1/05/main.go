package main

import (
	"fmt"
)

type water int

var x water

var y int

func main() {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
