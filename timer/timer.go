package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("...")
			t.Reset(2 * time.Second)
		}
	}
}
