package main

import "fmt"

func main() {
	for y := 0; y < 50; y++ {
		for x := 0; x < 100; x++ {
			y1 := (25.0 - float32(y)) / float32(50) * 3
			x1 := (float32(x) - 50.0) / 100.0 * 3
			tmp := (x1*x1 + y1*y1 - 1)
			if tmp*tmp*tmp-x1*x1*y1*y1*y1 < 0 {
				fmt.Printf("%c", '*')
			} else {
				fmt.Printf("%c", ' ')
			}
		}
		fmt.Println()
	}
}
