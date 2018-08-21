package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBuffer([]byte("hello, my name is domac"))

	br := bufio.NewReaderSize(buff, 20)
	pb, _ := br.Peek(5)
	fmt.Printf("%s\n", string(pb))

	pb2, _ := br.Peek(5)
	fmt.Printf("%s\n", string(pb2))

	pb3 := make([]byte, 5)
	br.Read(pb3)
	fmt.Printf("%s\n", string(pb3))

	br.Discard(2)

	pb4, _ := br.Peek(8)
	fmt.Printf("%s\n", string(pb4))

	pb5, _ := br.Peek(8)
	fmt.Printf("%s\n", string(pb5))

	size := br.Size() // buf size
	fmt.Printf("size: %d\n", size)
}
