package floatsort

import (
	"sort"
	"unsafe"
)

func FloatSort(a []float64) {
	sort.Float64s(a)
}

func FloatSort2(a []float64) {
	b := ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
}
