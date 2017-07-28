package stringCombineTest

import "testing"

var strs []string = []string{
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

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsJoin(strs)
	}
}

func BenchmarkStringsString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsString(strs)
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuffer(strs)
	}
}
func BenchmarkStringByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringByte(strs)
	}
}
