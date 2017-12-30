package main

import (
	"fmt"
)

func main() {
	var a uint16 = 0xCEFF //1100111011111111 - 52991
	fmt.Printf("%b - %d\n", a, a)

	a ^= 0xFF00
	fmt.Printf("%b - %d\n", a, a)

	x, y := -12, 25
	fmt.Println("a and b have same sign?", (x^y) >= 0)
}
