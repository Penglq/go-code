package testing

import (
	"testing"
	"go-code/change"
)

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		change.Builder()
	}
}

func BenchmarkSBuilder(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		change.SBuilder()
	}
}
