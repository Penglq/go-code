package main

import (
	"go-code/goroutine/work"
	"log"
	"time"
)

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Print(m.name)
	time.Sleep(time.Second)
}

func main() {

	names := []string{
		"steve",
		"bob",
		"mary",
		"therese",
		"jason",
	}

	p := work.New(2)

	for _, name := range names {
		np := namePrinter{
			name: name,
		}

		p.Add(&np)

	}
	p.Shutdown()
}
