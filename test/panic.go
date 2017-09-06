package main

import (
	"fmt"
	"reflect"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(reflect.TypeOf(r))
			fmt.Println(r)
		}
	}()
	panic([]int{11111})
}
