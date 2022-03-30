package main

import(
	"github.com\mthompson1998\learninggo-5\27_exercises12\01\dog"
	"fmt"
)

type canine struct {
	name string
	age int
}

func main() {
	fido := canine{
		name: "fido",
		age: dog.Years(10),
	}
	fmt.Println(fido)
}