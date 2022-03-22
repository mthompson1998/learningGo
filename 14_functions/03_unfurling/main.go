package main

import (
	"fmt"
)

// unfurling a slice of int //

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The total is", sum)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for the item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("the total is", sum)
	return sum
}