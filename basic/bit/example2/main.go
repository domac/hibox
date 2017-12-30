package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for x := 0; x < 100; x++ {
		//50以内的随机数
		num := rand.Int31n(50)
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}
}
