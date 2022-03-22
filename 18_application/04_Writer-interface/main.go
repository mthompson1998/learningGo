package main

import (
	"fmt"
	"os"
)
func main() {
	fmt.Println("hello")
	fmt.Fprintln(os.Stdout, "Hello")
}