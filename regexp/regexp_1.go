package main

import (
	"regexp"
	"fmt"
)

func main() {
	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)

	str := "2017-02-02 10:10:02"

	res := reg.MatchString(str)

	fmt.Print(res)
}
