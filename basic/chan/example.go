package main

import (
	"fmt"
	"time"
)

func main() {

	inChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			inChan <- i
		}
		close(inChan)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer

	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			println("--------------")
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-inChan:
			if !ok {
				fmt.Println("End .")
				return
			}
			fmt.Printf("Received : %v\n", e)
		case <-timer.C:
			fmt.Println("Timeout")
		}
	}
}
