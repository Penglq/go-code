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
			time.Sleep(4 * time.Second)
			fmt.Print(j)
			wg.Done()
		}(i)
	}
	//wg.Add(1)
	//go func() {
	//	log.Print("a")
	//	wg.Done()
	//}()
	wg.Wait()

}
