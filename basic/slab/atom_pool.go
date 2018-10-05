package slab

import (
	"reflect"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type AtomPool struct {
	classes []atomClass
	minSize int
	maxSize int
}

func NewAtomPool(minSize, maxSize, factor, pageSize int) *AtomPool {
	pool := &AtomPool{
		classes: make([]atomClass, 0, 10),
		minSize: minSize,
		maxSize: maxSize,
	}

	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		chunkLen := pageSize / chunkSize
		c := atomClass{
			size:   chunkSize,
			page:   make([]byte, pageSize),
			chunks: make([]chunk, chunkLen),
			head:   1 << 32,
		}

		for i := 0; i < len(c.chunks); i++ {
			index := i + 1
			chk := &c.chunks[i]
			chk.mem = c.page[i*chunkSize : (i+1)*chunkSize : (i+1)*chunkSize]
			if i < len(c.chunks)-1 {
				chk.next = uint64(index+1) << 32
			} else {
				c.pageBegin = uintptr(unsafe.Pointer(&c.page[0]))
				c.pageEnd = uintptr(unsafe.Pointer(&chk.mem[0]))
			}
		}
		pool.classes = append(pool.classes, c)
	}

	return pool
}

func (pool *AtomPool) Alloc(size int) []byte {
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

func (pool *AtomPool) Free(mem []byte) {
	size := cap(mem)
	for i := 0; i < len(pool.classes); i++ {
		if pool.classes[i].size == size {
			pool.classes[i].Push(mem)
			break
		}
	}
}

type atomClass struct {
	size      int
	page      []byte
	pageBegin uintptr
	pageEnd   uintptr
	chunks    []chunk
	head      uint64
}

type chunk struct {
	mem  []byte
	aba  uint32
	next uint64
}

func (c *atomClass) Push(mem []byte) {
	ptr := (*reflect.SliceHeader)(unsafe.Pointer(&mem)).Data
	if c.pageBegin <= ptr && ptr <= c.pageEnd {
		i := (ptr - c.pageBegin) / uintptr(c.size)
		chk := &c.chunks[i]
		if chk.next != 0 {
			panic("slab.AtomPool: Double Free")
		}
		chk.aba++
		new := uint64(i+1)<<32 + uint64(chk.aba)
		for {
			old := atomic.LoadUint64(&c.head)
			atomic.StoreUint64(&chk.next, old)
			if atomic.CompareAndSwapUint64(&c.head, old, new) {
				break
			}
			runtime.Gosched()
		}
	}
}

func (c *atomClass) Pop() []byte {
	for {
		old := atomic.LoadUint64(&c.head)
		if old == 0 {
			return nil
		}
		chk := &c.chunks[old>>32-1]
		nxt := atomic.LoadUint64(&chk.next)
		if atomic.CompareAndSwapUint64(&c.head, old, nxt) {
			atomic.StoreUint64(&chk.next, 0)
			return chk.mem
		}
		runtime.Gosched()
	}
}
