package slab

import (
	"reflect"
	"sync"
	"unsafe"
)

type LockPool struct {
	classes []lockClass
	minSize int
	maxSize int
}

func NewLockPool(minSize, maxSize, factor, pageSize int) *LockPool {
	n := 0
	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		n++
	}

	pool := &LockPool{
		classes: make([]lockClass, n),
		minSize: minSize,
		maxSize: maxSize,
	}

	n = 0
	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		chunkLen := pageSize / chunkSize
		c := &pool.classes[n]
		c.size = chunkSize
		c.page = make([]byte, pageSize)
		c.chunks = make([][]byte, chunkLen)
		c.head = 0
		c.tail = chunkLen - 1

		for i := 0; i < len(c.chunks); i++ {
			c.chunks[i] = c.page[i*chunkSize : (i+1)*chunkSize : (i+1)*chunkSize]
			if i == len(c.chunks)-1 {
				c.pageBegin = uintptr(unsafe.Pointer(&c.page[0]))
				c.pageEnd = uintptr(unsafe.Pointer(&c.chunks[i][0]))
			}
		}
		n++
	}
	return pool
}

func (pool *LockPool) Alloc(size int) []byte {
	if size <= pool.maxSize {
		for i := 0; i < len(pool.classes); i++ {
			if pool.classes[i].size >= size {
				mem := pool.classes[i].Pop()
				if mem != nil {
					return mem[:size]
				}
				break
			}
		}
	}
	return make([]byte, size)
}

func (pool *LockPool) Free(mem []byte) {
	size := cap(mem)
	for i := 0; i < len(pool.classes); i++ {
		if pool.classes[i].size == size {
			pool.classes[i].Push(mem)
			break
		}
	}
}

type lockClass struct {
	sync.Mutex
	size int

	page      []byte
	pageBegin uintptr
	pageEnd   uintptr

	chunks [][]byte

	head int
	tail int
}

func (l *lockClass) Push(mem []byte) {
	ptr := (*reflect.SliceHeader)(unsafe.Pointer(&mem)).Data

	if l.pageBegin <= ptr && ptr <= l.pageEnd {
		l.Lock()
		l.tail++
		n := l.tail % len(l.chunks)
		if l.chunks[n] != nil {
			panic("slab lockclass: double free")
		}
		l.chunks[n] = mem
		l.Unlock()

	}
}

func (l *lockClass) Pop() []byte {
	var mem []byte
	l.Lock()
	if l.head <= l.tail {
		n := l.head % len(l.chunks)
		mem = l.chunks[n]
		l.chunks[n] = nil
		l.head++
	}
	l.Unlock()
	return mem
}
