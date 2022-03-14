package main

// marshalling data enables you to send data outside of the go program
import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string // must be capitalized 
	Last string
	Age int
}

func main() {
	p1 := person {
		First: "James",
		Last: "Bond",
		Age: 32,
	}
	p2 := person {
		First: "Miss",
		Last: "Moneypenny",
		Age: 27,
	}

	people := []person{p1,p2}

	fmt.Println(people)

	bs, err := json.Marshal(people) // bs is byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}