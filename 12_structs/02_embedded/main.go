package main

import (
	"fmt"
)
// Embedding 'person' type within type secretAgent
type person struct{
	first string
	last string
	age int
}

type secretAgent struct{
	person 
	ltk bool
	//ltk = lisence to kill
}

func main() {

		sa1 := secretAgent{
			person: person{
				first: "James",
				last: "Bond",
				age: 32,
			},
			ltk: true,
		}

		p2 := person{
			first: "Miss",
			last: "Moneypeny",
			age: 27,
		}
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}