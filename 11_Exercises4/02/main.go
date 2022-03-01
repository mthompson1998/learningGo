package main

import (
	"fmt"
)

/*
Using a composite literal
	- create a slice of type INT
	- assing 10 values
Range over the array and print the values out.
Using format printing
	print out the type of the array
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)

}
