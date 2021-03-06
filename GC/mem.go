package main

import (
	"fmt"
	"runtime"
	"time"
)

type Garbage struct{ a int }

func notify(f *Garbage) {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)

	fmt.Println("Last GC was:", stats.LastGC)

	go ProduceFinalizedGarbage()
}

func ProduceFinalizedGarbage() {
	x := &Garbage{}
	runtime.SetFinalizer(x, notify)
}

func main() {
	go ProduceFinalizedGarbage()

	for {
		runtime.GC()
		time.Sleep(30 * time.Second) // Give GC time to run
	}
}
