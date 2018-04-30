package floatsort

import (
	"testing"
)

func BenchmarkFloatSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
		FloatSort(a)
	}
}

func BenchmarkFloatSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
		FloatSort2(a)
	}
}
