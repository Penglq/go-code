package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000000)+5000000000)
}

func main() {
	pool := make([][]byte, 20)

	var m runtime.MemStats
	makes := 0
	for {
		b := makeBuffer()
		makes += 1
		i := rand.Intn(len(pool))
		pool[i] = b

		time.Sleep(time.Millisecond/100)

		bytes := 0

		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("程序向应用程序申请的内存:%d,%d,堆上目前分配的内存:%d,堆上目前没有使用的内存:%d,回收到操作系统的内存:%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes)
	}
}
