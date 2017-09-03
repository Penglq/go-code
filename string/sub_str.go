package main

import "fmt"

func SubStr(start int8, end int8, str []rune) string {
	return string(str[start:end])
}

func main() {
	str := "2017-01-02 10:20:20"

	fmt.Println(SubStr(0,10, []rune(str)))
}
