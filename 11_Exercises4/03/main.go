package main

import (
	"fmt"
)

/*
print out 3 outputs with slices of the array at different points
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(x)
	fmt.Println(x[3:])
	fmt.Println(x[2:6])
	fmt.Println(x[:5])

}
