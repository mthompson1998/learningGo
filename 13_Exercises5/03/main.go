package main

import (
	"fmt"
)

type vehicle struct{
	doors string
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
		doors: "two",
		color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
		doors: "four",
		color: "grey",
		},
		luxury: false,
	}

	fmt.Println(t.doors, t.color, t.fourWheel)
	fmt.Println(s.doors, s.color, s.luxury)

}