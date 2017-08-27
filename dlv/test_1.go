package main

import (
	"log"
	"runtime"
	"runtime/debug"
	"fmt"
)

func main() {
	var a int = 1
	fmt.Print(a)
	runtime.GC()

	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	//s = a + b
	log.Printf("%#v", stats.Alloc)

	st := debug.GCStats{}
	debug.ReadGCStats(&st)
	log.Printf("%#v", st)
}
