package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v M", m.Alloc/1024/1024)
	fmt.Printf("\tHeap分配 = %v M", m.HeapAlloc/1024/1024)
	fmt.Printf("\t总分配 = %v M", m.TotalAlloc/1024/1024)
	fmt.Printf("\tGC次数 = %v\n", m.NumGC)
}

func StringMap() {
	PrintMemUsage()
	y := make(map[string]string)
	for i := 0; i < 50000000; i++ {
		test := strconv.Itoa(i)
		y[test] = test
		if (i+1)%100000 == 0 {
			fmt.Println("inserted ", i+1, " items")
			PrintMemUsage()
		}
	}
}

func IntMap() {
	PrintMemUsage()
	y := make(map[int]string)
	for i := 0; i < 50000000; i++ {
		test := strconv.Itoa(i)
		y[i] = test
		if (i+1)%100000 == 0 {
			fmt.Println("inserted ", i+1, " items")
			PrintMemUsage()
		}
	}
}

func main() {
	IntMap()
}
