package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)

	<-out
	out <- 2
	//fmt.Println(a)
	//go f1(out)
}