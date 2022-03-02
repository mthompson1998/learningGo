package main

import (
	"fmt"
)
// store the values of type person in the map
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
			"chocolate",
			"brownie",
			"shorbet",
		},

	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m{
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.flavor{
			fmt.Println(i, val)
		}
		fmt.Println("--------------")
	}
}