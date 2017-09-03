package main

import (
	"time"
	"fmt"
)

func main() {
	workTimer := time.Tick(2 * time.Second)
	for {
		select {
		case <-workTimer:
			//do()
			fmt.Print(1)
		default:
			fmt.Print("  .  ")
			time.Sleep(3 * time.Second)
		}
	}
}
