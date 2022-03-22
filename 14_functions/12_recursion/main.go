package main

import(
	"fmt"
)

// can have negative impacts on memory, not always useful
// recusion is when a function calls itself

//example : factorial

// often easier to do a loop

func main() {
	fmt.Println(4*3*2*1)
	n := factorial(6)
	fmt.Println(n)

	s1 := loopFact(4)
	fmt.Println(s1)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// rewrite with loop

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}