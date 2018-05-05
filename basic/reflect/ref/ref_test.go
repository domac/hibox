package ref

import (
	"testing"
)

func BenchmarkFiltName(b *testing.B) {

	for i := 0; i < b.N; i++ {
		u := &User{}
		u.Name = "dom"
		u.Age = i
		FiltName(u, "can you give me some beer?")
	}

}

func BenchmarkFiltNameWithCache(b *testing.B) {

	for i := 0; i < b.N; i++ {
		u := &User{}
		u.Name = "dom"
		u.Age = i
		FiltNameWithReuseOffset(u, "can you give me some beer?")
	}

}
