package main

import "fmt"

func main() {
	a := 1
	defer fmt.Println(a+3)
	a = 2
	fmt.Println(a)
}