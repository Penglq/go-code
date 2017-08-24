package main

import "fmt"

func Foo(x interface{}) {
	fmt.Printf("%v", x)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}
