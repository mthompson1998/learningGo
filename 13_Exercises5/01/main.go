package main

import (
	"fmt"
)

type person struct{
	first string
	last string
	flavor []string
}
func main() {
	p1 := person{
		first: "Mike",
		last: "Thompson",
		flavor: []string{
			"vanilla",
			"cookie dough",
			"lemon",
		},
	}
	p2 := person{
		first: "Abigail",
		last: "Lippincott",
		flavor: []string{
			"Chocolate",
			"brownie",
			"sherbet",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.flavor{
		fmt.Println(i,v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.flavor{
		fmt.Println(i,v)
	}

}