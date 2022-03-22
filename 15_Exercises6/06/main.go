package main

import (
	"fmt"
)

func main() {
	x := 42
	func() {
		fmt.Println("the meaning of life is", x)
	}()

	func(){
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("Anonymous functions!")
}