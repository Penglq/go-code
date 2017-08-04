package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
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
		atomic.AddInt64(&counter, 1)

		runtime.Gosched()
	}
}
