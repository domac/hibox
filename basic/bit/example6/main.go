package main

import (
	"fmt"
)

func main() {
	var a byte = 0xAB
	fmt.Printf("%b - %d\n", a, a)

	var b byte = 0x0F
	fmt.Printf("%b\n", b)

	a &^= b
	fmt.Printf("%b\n", a)
}
