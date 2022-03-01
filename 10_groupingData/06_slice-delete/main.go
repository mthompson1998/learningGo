package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6, 7, 8, 9, 10)
	fmt.Println(x)

	y := []int{100, 123, 122}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
