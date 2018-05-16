package main

import "fmt"

func test(a int) func(i int) int {
	return func(i int) int {
		a = a + i
		return a
	}
}
func main() {
	f := test(1)
	a := f(2)
	fmt.Println(a)
	b := f(3)
	fmt.Println(b)
}
