package main

import (
	"flag"
	"log"
)

func main() {
	var s string
	// go run stringVar.go -name=aa
	flag.StringVar(&s, "name", "12", "注释")
	flag.Parse()
	log.Print(s)
}
