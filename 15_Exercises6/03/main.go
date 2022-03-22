package main

import (
	"fmt"
)

func main() {
defer foo()
	bar()
}

func foo() {
	fmt.Println("defer")
}

func bar() {
	fmt.Println("send to first")
}