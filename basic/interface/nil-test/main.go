package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}
func f(out io.Writer) {
	if out != nil {
		fmt.Println("surprise!")
	}
}
