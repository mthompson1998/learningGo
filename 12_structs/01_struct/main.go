package main

import (
	"fmt"
)
// allow you to aggregate values of different types
// create VALUES of TYPE
type person struct{
	first string
	last string
	age int
}

func main() {

		p1 := person{
			first: "James",
			last: "Bond",
			age: 33,
		}

		p2 := person{
			first: "Miss",
			last: "Moneypeny",
			age: 27,
		}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}