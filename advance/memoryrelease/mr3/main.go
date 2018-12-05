package main

import (
	"runtime"
	"time"
)

type Node struct {
	next     *Node
	playload [64]byte
}

func f() {
	curr := new(Node)
	for i := 0; i < 10000000; i++ {
		curr.next = new(Node)
		curr = curr.next
	}
}

func main() {
	f()
	time.Sleep(time.Second * 300) //在这里查看进程内存使用
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	println(memStats.HeapInuse)
}
