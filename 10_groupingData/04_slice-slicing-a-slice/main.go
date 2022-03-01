package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[2:])
		// [3, 4, 5] (spot 3 to the end)
	fmt.Println(x[1:4])
		// [2, 3, 4, 5] (from spot 2 to the end/spot 4)
}