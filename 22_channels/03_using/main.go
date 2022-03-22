package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	//send
	go foo(c)

	//receive
	bar(c) // pull of the "go" before statement to hold output until foo is finished

	fmt.Println("output")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
