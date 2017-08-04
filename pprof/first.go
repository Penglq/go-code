package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		i := 1
		if i == 2 {

		}
	}
}
