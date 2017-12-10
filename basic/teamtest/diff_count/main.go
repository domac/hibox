package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"
)

var statsMap = make(map[uint64]uint64, 4096)
var nameMap = make(map[uint64]string, 256)

func bytes64(data []byte) uint64 {
	var v uint64 = 14695981039346656037
	for _, c := range data {
		v = (v ^ uint64(c)) * 1099511628211
	}
	return v
}

func hashFunc(data []byte) uint64 {
	h := bytes64(data)
	if _, ok := nameMap[h]; !ok {
		nameMap[h] = string(data)
	}
	return h
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func parsebyteToUint(b []byte) uint64 {
	var n uint64
	i := 0
	for ; i < len(b); i++ {
		var v byte
		d := b[i]
		v = d - '0'
		n *= uint64(10)
		n1 := n + uint64(v)
		n = n1
	}
	return n
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
			hashVal := hashFunc(b[0:idx])
			statsMap[hashVal] += parsebyteToUint(b[idx+1:])
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
		fmt.Printf("[%s]:%d\n", nameMap[k], v)
	}
}
