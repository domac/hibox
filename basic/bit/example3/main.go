package main

import (
	"fmt"
)

//Using OR is quite useful when doing bit masking techniques to set arbitrary bits for a given integer value.
//For instance, we can expand the previous program to set more bits in the value stored in variable a.
func main() {
	var a uint8 = 0
	//a = a | 196
	a |= 196
	fmt.Printf("%b - %d\n", a, a)
	a |= 1
	fmt.Printf("%b - %d\n", a, a)

}
