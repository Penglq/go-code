package main

import (
	"sync"
	"fmt"
)
// GOMAXPROCS=1 go run echoSort.go
// 结果: 9012345678
func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			fmt.Println(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
