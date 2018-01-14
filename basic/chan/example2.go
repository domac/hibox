package main

import (
	"fmt"
)

func main() {

	done := make(chan int, 1)

	go func() {
		fmt.Println("hello world")
		done <- 1
	}()

	<-done
	println("done")
}
