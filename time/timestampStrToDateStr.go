package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "1504084457"
	// string到int
	int, err := strconv.Atoi(str)
	// string到int64
	int64, err := strconv.ParseInt(str, 10, 64)
	// int到string
	s1 := strconv.Itoa(int)
	// int64到string
	s2 := strconv.FormatInt(int64, 10)

	fmt.Print(s2, s1, err)
}
