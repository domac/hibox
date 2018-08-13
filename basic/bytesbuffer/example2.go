package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBuffer([]byte("hello, my name is domac"))

	wordCount := 5

	block := make([]byte, wordCount)

	for buff.Len() > 0 {
		block = buff.Next(5)

		fmt.Printf("read data : %s\n", block)
	}

	block = block[0:]
}
