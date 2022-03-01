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

	// add element to map
	m["todd"] = 33

	if v, ok := m["Mike"]; ok {
		fmt.Println(v)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
