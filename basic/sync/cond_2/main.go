package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)

//条件量需要依赖一个互斥锁
var cond = sync.NewCond(locker)

//当主gouroutine执行完任务后，通过BroadCast广播信号。
//处于cond.Wait状态的所有gouroutine收到信号后将全部被唤醒并往下执行。
//需要注意的是，从gouroutine执行完任务后，需要通过cond.L.Unlock释放锁， 否则其它被唤醒的gouroutine将没法继续执行。
func condTest(i int) {

	//下面两对加锁-解锁方法效果一样
	//cond.L.Lock()
	//defer cond.L.Unlock()
	locker.Lock()
	defer locker.Unlock()

	//等待广播通知
	//Wait函数会在返回前，重新对关联的互斥锁进行锁上， 所以为了不影响其他并发程序的锁定操作，需要在程序处理完的时候，进行对关联的互斥锁进行解锁：defer locker.Unlock()
	cond.Wait()
	fmt.Println(i)
	time.Sleep(1 * time.Second)

}

func main() {
	for i := 0; i < 100; i++ {
		go condTest(i)
	}
	fmt.Println("start main now")
	cond.Broadcast()
	fmt.Println("broadcast finish")
	time.Sleep(100 * time.Second)
}
