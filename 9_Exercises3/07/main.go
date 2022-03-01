package main

import (
	"fmt"
)

// create an If Statement with else if

func main() {
	t := 80
	if t == 75 {
		fmt.Println("its nice outside!")
	} else if t == 80 {
		fmt.Println("wear your sunscreen!", t)
	} else {
		fmt.Println("check the weather")
	}
}