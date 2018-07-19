package main

import (
	"sync"
	"sync/atomic"
)

var g_RedisLock = make(map[string]*MidRedisLock)

var glockLock sync.RWMutex

func GetRedisLock(mid string) *MidRedisLock {
	glockLock.RLock()
	defer glockLock.RUnlock()
	return g_RedisLock[mid]
}

type MidRedisLock struct {
	mid   string
	sema  uint32
	MLock *sync.RWMutex
}

func NewRedisLock(mid string) *MidRedisLock {
	return &MidRedisLock{
		mid:   mid,
		sema:  uint32(0),
		MLock: new(sync.RWMutex),
	}
}

type TestModel struct {
	TMidRedisLock *MidRedisLock
}

func (t *TestModel) TaskSeqLock(mid string) {
	t.TMidRedisLock = GetRedisLock(mid)

	if t.TMidRedisLock == nil {
		glockLock.Lock()
		if t.TMidRedisLock == nil {
			midRedisLock := NewRedisLock(mid)
			t.TMidRedisLock = midRedisLock
			g_RedisLock[mid] = midRedisLock
		}
		glockLock.Unlock()
	}

	if atomic.CompareAndSwapUint32(&t.TMidRedisLock.sema, 0, 1) {
		t.TMidRedisLock.MLock.Lock()
	}
}

func main() {
}
