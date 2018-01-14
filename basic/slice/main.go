package main

import (
	"fmt"
)

func test1() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}

func test2() {
	s := []int{5, 7, 9}
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}

func main() {
	test1()
	test2()
}
