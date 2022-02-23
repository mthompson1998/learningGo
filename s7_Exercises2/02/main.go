package main

import (
	"fmt"
)

/*
Use the following operators, write expressions and assign their values to variables
	g. ==
	h. <=
	i. >=
	j. !=
	k. <
	l. ?
Now print each variable
*/

func main() {

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)

	fmt.Println(a, b, c, d, e)

}
