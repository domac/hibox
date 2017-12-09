package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var splitB = []byte(":")

var statsMap = make(map[string]int64, 4096)

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
			//readCh <- bytes.Split(b, splitB)
			bs := bytes.Split(b, splitB)
			v, _ := strconv.ParseInt(string(bs[1]), 10, 0)
			statsMap[string(bs[0])] += v
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
	if dataFile != "" {
		readAndHandleDataFile(dataFile)
	}
	elapsed := time.Now().Sub(start)
	log.Printf("total elapsed time: %f seconds", elapsed.Seconds())

	for k, v := range statsMap {
		fmt.Printf("[%s]:%d\n", k, v)
	}
}
