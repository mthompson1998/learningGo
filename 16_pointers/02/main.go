package main

import (
	"fmt"
)

/*
When to use a pointer?
	- when you need to change a value
	- when you need to change the data at a location
	- when you dont want to pass around a lot of data
Every thing in Go is pass by value
*/
func main() {
	x := 0
	fmt.Println(&x)
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)

}
func foo(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}
