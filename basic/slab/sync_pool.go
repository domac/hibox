package slab

import (
	"sync"
)

type SyncPool struct {
	classes []SyncClass
	minSize int
	maxSize int
}

func NewSyncPool(minSize, maxSize, factor int) *SyncPool {
	n := 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		n++
	}
	pool := &SyncPool{
		classes: make([]SyncClass, n),
		minSize: minSize,
		maxSize: maxSize,
	}

	n = 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		pool.classes[n].size = chunkSize
		pool.classes[n].chunks.New = func(size int) func() interface{} {
			return func() interface{} {
				buf := make([]byte, size)
				return &buf
			}
		}(chunkSize)
		n++
	}
	return pool
}

func (pool *SyncPool) Alloc(size int) []byte {
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

func (pool *SyncPool) Free(mem []byte) {
	if size := cap(mem); size <= pool.maxSize {
		for i := 0; i < len(pool.classes); i++ {
			if pool.classes[i].size >= size {
				pool.classes[i].Push(mem)
				return
			}
		}
	}
}

type SyncClass struct {
	size   int
	chunks sync.Pool
}

func (s *SyncClass) Push(mem []byte) {
	s.chunks.Put(&mem)
}

func (s *SyncClass) Pop() []byte {
	mem := s.chunks.Get().(*[]byte)
	return *mem
}
