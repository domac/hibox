package main

import (
	"fmt"
)

//The AND operator has the nice side effect of selectively clearing bits of an integer value to zero.
//For instance,
//we can use the & operator to clear (set to zero) the last 4 least significant bits (LSB) to all zeros

func main() {
	var x uint8 = 0xAC //10101100 //172
	fmt.Printf("%b - %d\n", x, x)
	var y uint8 = 0xF0
	fmt.Printf("%b - %d\n", y, y)

	x1 := x & y
	fmt.Printf("%b - %d\n", x1, x1)

	x &= y
	fmt.Printf("%b - %d\n", x, x)

}
