package main

import (
	"fmt"
)

// create an anonymous struct

func main() {
	book := struct{
		title string
		author string
		ID int
	}{
		author: "Ian Fleming",
		title: "Casino Royale",
		ID: 001,
	}
	fmt.Println(book)
}