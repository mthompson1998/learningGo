package main

import (
	"fmt"
)

var a int

// variable called a with TYPE integer

type hotdog int

// new TYPE of hotdog with UNDERLYING TYPE integer

var b hotdog

// varaible called b with type hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

/*
OUTPUT
$ go run main.go
42
int
43
main.hotdog
*/
