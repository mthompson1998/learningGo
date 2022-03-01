package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Moneypenny"])

	v, ok := m["Mike"]
	fmt.Println(v)
	fmt.Println(ok)

	// delete element to map
	delete(m, "James")
	fmt.Println(m)
	//use idom to varify the delete
	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Moneypenny")
	}
}
