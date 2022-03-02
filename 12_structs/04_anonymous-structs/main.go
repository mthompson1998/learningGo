package main

import (
	"fmt"
)

// anonymous struct. declared in p1, not defined as "person"
// fields and types defined above the values
// used to keep code lean (no global scope)

func main() {
	p1 := struct{
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}
	fmt.Println(p1)
}