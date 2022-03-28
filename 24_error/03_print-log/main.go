package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("when os.Exit() is called, deferred functions don't run")
}

// the fatal functions call os.Exit(1) after writing the log message

// fatalln is equivalent to Println() folowed by a call to os.Exit(1)
