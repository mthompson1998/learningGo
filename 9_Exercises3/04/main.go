package main

import (
	"fmt"
)

// do the same thing but in a different for loop style

func main() {
	bd := 1998
	for {
		if bd > 2022{
			break
		}
		fmt.Println(bd)
		bd++
	}
}