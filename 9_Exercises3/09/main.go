package main

import (
	"fmt"
)

// create switch statement

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains")
	case "swimming":
		fmt.Println("go to the pool")
	case "surfing":
		fmt.Println("go to the ocean")
	case "wingsuiting":
		fmt.Println("go to the hospital")
	}
}