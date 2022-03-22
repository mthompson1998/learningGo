package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}
func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

	bs := []byte(s) // byte is an alias for uint8
	fmt.Printf("%T\n", s) // string
	fmt.Printf("%T\n", bs) //uint8

	var people []person

	err := json.Unmarshal(bs, &people) // pointed to people
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
