package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var mutex = new(sync.RWMutex)

var sema uint32

func MyLock(i int) {
	mutex.Lock()
	println("----", i)
}

func MySafeLock(i int) {
	for {
		if atomic.LoadUint32(&sema) == 0 {
			break
		}
	}

	if atomic.CompareAndSwapUint32(&sema, 0, 1) {
		mutex.Lock()
		println("----", i)
	}
}

func MyUnlock() {
	mutex.Unlock()
	atomic.StoreUint32(&sema, 0)
}

func main() {

	MySafeLock(1)
	MySafeLock(2)
	println("00001")

	MyUnlock()
	MyUnlock()

	time.Sleep(5 * time.Second)
}
