package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p := Person{Name:"a"}

	fmt.Printf("%#v", p == Person{})
}
