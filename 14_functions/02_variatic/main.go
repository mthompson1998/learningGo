package main

import (
	"fmt"
)

//variatic parameter functions

/*
func (r receiver) identifier(parameter(s)) (return(s)){
	--code--
}
*/


func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the total is,", x)
}

func sum(x ...int) int{
	fmt.Println(x)
	fmt.Println("%T\n", x)
	
	sum := 0
	for i, v := range x{
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the toal which is,", sum)
	
	}
	fmt.Println("the total is,", sum)
	return sum
}

