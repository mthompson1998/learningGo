package main

import (
	"fmt"
)


func main() {
	x := []int{5, 4, 3, 5, 2, 7}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	// range slice x and display value and index
	for i, v := range x {
		fmt.Println(i, v)
	}
}