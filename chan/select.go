package main

import "fmt"

func main() {
	//runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	string_chan <- "hello"
	int_chan <- 1
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		//fmt.Println(value)
		panic(value)
	}
}
