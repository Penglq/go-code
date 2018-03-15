package main

import (
	"fmt"
	"sync"
)

var (
	val map[string]int
	wg  sync.WaitGroup
)

func main() {

	val = make(map[string]int)
	val["a"] = 0

	wg.Add(2)
	go incVal()
	go incVal()

	wg.Wait()
	fmt.Println(val)
}

func incVal() {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		val["a"]++

		//runtime.Gosched()
	}
}
