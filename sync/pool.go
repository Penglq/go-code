package main

import (
	"sync"
	"log"
)

func main() {
	var p sync.Pool
	if p.Get() != nil {
		log.Print("expected empty")
	}
	p.Put("a")
	g := p.Get()
	h := p.Get()
	log.Print(g)
	log.Print(h)
}
