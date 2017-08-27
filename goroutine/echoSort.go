package main

import (
	"sync"
)

// GOMAXPROCS=1 go run echoSort.go
// 结果: 9012345678
func main() {
	//runtime.GOMAXPROCS(1)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			//time.Sleep(time.Duration(j) * time.Second)
			for n := 0; n < 10000000000; n++ {

			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}
