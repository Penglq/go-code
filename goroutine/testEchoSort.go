package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// GOMAXPROCS=1 go run echoSort.go
// 结果: 9012345678
func main() {
	runtime.GOMAXPROCS(1)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			time.Sleep(time.Duration(1) * time.Second)
			//for n := 0; n < 1000000000; n++ {
			//
			//}
			fmt.Print(j)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
