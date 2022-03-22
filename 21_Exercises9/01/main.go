package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
		}()

	go func() {
		fmt.Println("hello from thing two")
		wg.Done()
		}()
		wg.Wait()

		fmt.Println("about to exit")
		fmt.Println("begin cpu", runtime.NumCPU())
		fmt.Println("begin gs", runtime.NumGoroutine())
	}