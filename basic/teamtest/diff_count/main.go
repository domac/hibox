package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

var statsMap = make(map[uint64]uint64, 1024)
var nameMap = make(map[uint64]string, 1024)

//自定义哈希函数
func hashFunc(data []byte) uint64 {
	var h uint64 = 14695981039346656037
	for _, c := range data {
		h = (h ^ uint64(c)) * 20171111
	}

	if _, ok := nameMap[h]; !ok {
		nameMap[h] = string(data)
	}
	return h
}

//把字符数组转化为无符号整型
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

//处理数据文件
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
			idx := bytes.IndexByte(b, ':') //分隔符所在索引位置
			hashVal := hashFunc(b[0:idx])  //计算哈希值
			statsMap[hashVal] += parsebyteToUint(b[idx+1:])
		}
	}
	log.Printf("read %s completed !\n", filepath)
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
