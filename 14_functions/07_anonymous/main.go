package main

import (
	"fmt"
)


func foo() {
	fmt.Println("foo ran")
}


func main() {
	foo()

	func() {
		fmt.Println("anonymous func ran")
	}()

	func(x int) {
		fmt.Println("the meaning of life:", x)
	}(42)

	func(y int) {
		fmt.Println("Big Brother is watching,", y)
	}(1984)
}