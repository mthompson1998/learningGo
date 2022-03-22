package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int64

	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end count:", counter)
}