package main

import (
	"fmt"
)

func test1() {

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func test2() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func test3() {
	for i := 0; i < 10; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func main() {
	test3()
}
