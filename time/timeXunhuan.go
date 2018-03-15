package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	t2 := time.Now()
	var i=1
	for ; t2.Before(t1); i++ {

	}
	fmt.Print(i)
}
