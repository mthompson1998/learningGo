package main

import (
	"fmt"
)

type person struct{
	first string
	last string
}

type secretAgent struct{
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent {
		person: person {
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent {
		person: person {
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()
}