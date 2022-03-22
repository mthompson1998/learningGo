package main

import (
	"fmt"
)

// pointers point to the address where the memory is stored
func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	*b = 43 // assign the value stored at the address of b to 43. change from 42 to 43
	fmt.Println(a)

}
