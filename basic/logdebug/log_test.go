package main

import (
	"strings"
	"testing"
)

var a = strings.Repeat("abcde", 1024)

func BenchmarkDebug(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Debug(a)
	}
}
