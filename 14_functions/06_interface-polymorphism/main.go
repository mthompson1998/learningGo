package main

import (
	"fmt"
)
// Interfaces allow a value to be of more than one type //

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func main() {
	sa1 := secretAgent {
		person: person {
			"james",
			"bond",
		},
		ltk: true,
	}
	sa2 := secretAgent {
		person: person {
			"miss",
			"moneypenny",
		},
		ltk: false,
	}
	p1 := person{
		first: "dr.",
		last: "yes",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)

	sa1.speak()
	sa2.speak()
}