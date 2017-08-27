package main

import (
	"os"
	"os/signal"

	"log"
)

func main() {
	s := make(chan os.Signal, 1)

	signal.Notify(s)

	log.Print("get signal: ", <-s)
	os.Exit(-1)
}
