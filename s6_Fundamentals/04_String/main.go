package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
		// print the UTF value of each bit (byte?)
	}
}
