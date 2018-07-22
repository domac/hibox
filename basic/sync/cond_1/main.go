package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	locker := new(sync.RWMutex)
	cond := sync.NewCond(locker)
	done := false

	cond.L.Lock()

	go func() {
		time.Sleep(1 * time.Second)
		done = true
		cond.Signal()
	}()

	//这里当主goroutine进入cond.Wait的时候，就会进入等待
	//当从goroutine发出信号之后，主goroutine才会继续往下面走。
	if !done {
		cond.Wait()
	}

	fmt.Println("Task done : ", done)
}
