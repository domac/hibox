package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buff := bytes.NewBuffer([]byte("hello, my name is domac"))
	block := make([]byte, 5)

	for buff.Len() > 0 {
		n, err := buff.Read(block)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("read err: %s\n", err)
		} else {
			fmt.Printf("read bytes count :%d = %s\n", n, block)
		}
	}
	fmt.Println("read finish")

}
