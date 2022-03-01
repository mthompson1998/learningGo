package main

import (
	"fmt"
)

/*
Map Definition:
Built in data structure that associates values of one type (the key)
with values of another type (the element or value).

Maps hold references to an underlying data structure.
if you pass a map to a function that changes the contents
of the map, the changes will be visible in the caller
*/
func main() {
	// Make a map with the Key and the Value
	//Creating a composite literal with curlies
	// "key":	value
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

	// comma ok idiom
	// used to distinguish a missing entry from a zero value
	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

}
