package main

import (
	"fmt"
)

func main() {
	s := "我是一颗abcd"
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%#v,", s[i])
	}
	fmt.Println(len([]rune(s)))
	for _, v := range s {
		fmt.Printf("%c--%#v  ", v, v)
	}
}
