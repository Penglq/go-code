// 主程序本身算是一个协程
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
	}

	fmt.Print(runtime.NumGoroutine())
}
