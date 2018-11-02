package main

import (
	"fmt"
)

func _reverse(s *[]byte, idx *int, n *int) {
	if len(*s) > *idx {
		c := (*s)[*idx]
		*idx = *idx + 1
		_reverse(s, idx, n)
		(*s)[*n] = c
		*n = *n + 1
	}
}

func reverse(b *[]byte) {
	idx := 0
	n := 0
	_reverse(b, &idx, &n)
}

func main() {
	b := []byte{'d', 'o', 'm', 'a', 'c'}
	reverse(&b)
	fmt.Printf("result : %s \n", b)
}
