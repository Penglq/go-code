package main

import (
	"go-code/gotest/stringCombineTest"
	"log"
	"time"
)

type Testtest interface {
	Mytest(strs []string) string
}

type strString struct {
}

func (S strString) Mytest(strs []string) string {
	return stringCombineTest.StringsString(strs)
}

type strJoin struct {
}

func (s strJoin) Mytest(strs []string) string {
	return stringCombineTest.StringsJoin(strs)
}

type strButter struct {
}

func (s strButter) Mytest(strs []string) string {
	return stringCombineTest.StringBuffer(strs)
}

func tests(s []string, n int, t Testtest) {
	for i := 0; i < n; i++ {
		t.Mytest(s)
	}
}

func main() {
	strs := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}

	n := 10000000

	start1 := time.Now()
	tests(strs, n, strJoin{})
	log.Print("join时间", time.Since(start1))

	start2 := time.Now()
	tests(strs, n, strString{})
	log.Print("+时间", time.Since(start2))

	start3 := time.Now()
	tests(strs, n, strButter{})
	log.Print("buffer时间", time.Since(start3))

}
