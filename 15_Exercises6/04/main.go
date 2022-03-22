package main

import (
	"fmt"
)
type person struct {
	first string
	last string
	age int
}

//attache method to type person
func (p person) speak() {
	fmt.Println("hello, I am", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
	p1 := person{
		first: "Mike",
		last: "Thompson",
		age: 24,
	}
	fmt.Println(p1)
	p1.speak()
}