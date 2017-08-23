package main

import "fmt"

type Spec interface {
	Satisfy(s string) bool
}

type IsHave struct {
	S string
}

func (s IsHave) Satisfy(str string) bool {
	return str == s.S
}

func Exists(str []string, s Spec) bool {
	for _, v := range str {
		if s.Satisfy(v) {
			return true
		}
	}
	return false
}

func main() {
	s := []string{"ab", "bc", "c"}
	isHave := IsHave{S:"b"}

	fmt.Print(Exists(s, isHave))
}
