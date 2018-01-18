package main

import (
	"fmt"
	"sync"
)

//使用两个 goroutine 交替打印序列
//一个 goroutinue 打印数字， 另外一个goroutine打印字母
//最终效果如下 12AB34CD56EF78GH910IJ 。
func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	chanA := make(chan bool, 1)
	chanB := make(chan bool)

	//初始化
	chanA <- true

	//g1
	go func() {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < 9; i++ {
			<-chanA
			fmt.Printf("%d", a[i])
			fmt.Printf("%d", a[i+1])
			i++
			chanB <- true
		}
		wg.Done()
	}()

	//g2
	go func() {
		b := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		for i := 0; i < 9; i++ {
			<-chanB
			fmt.Printf("%s", b[i])
			fmt.Printf("%s", b[i+1])
			i++
			chanA <- true
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done !")
}
