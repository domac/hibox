package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"
)

var splitB = []byte(":")

var statsMap = make(map[string]uint64, 4096)

const intSize = 32 << (^uint(0) >> 63)
const maxUint64 = (1<<64 - 1)

var ErrRange = errors.New("value out of range")
var ErrSyntax = errors.New("invalid syntax")

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func parseUint(s string) (uint64, error) {
	var n uint64
	var err error
	i := 0
	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		v = d - '0'
		n *= uint64(10)
		n1 := n + uint64(v)
		n = n1
	}
	return n, err
}

func readAndHandleDataFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()
	s := bufio.NewScanner(f)

	for s.Scan() {
		if b := s.Bytes(); b != nil {
			idx := bytes.IndexByte(b, ':')
			k := b[0:idx]
			v := b[idx+1:]
			count, _ := parseUint(bytesToString(v))
			statsMap[string(k)] += count
		}
	}
	log.Printf("read  %s completed !\n", filepath)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please tell me the data file")
	}
	dataFile := args[1]

	start := time.Now()
	log.Println("read start")
	if dataFile != "" {
		readAndHandleDataFile(dataFile)
	}
	elapsed := time.Now().Sub(start)
	log.Printf("total elapsed time: %f seconds", elapsed.Seconds())

	for k, v := range statsMap {
		fmt.Printf("[%s]:%d\n", k, v)
	}
}
