package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Print(counter)
}

func incCounter() {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		val := counter

		runtime.Gosched()

		val++

		counter = val
	}
}
