package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 10; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("i:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
