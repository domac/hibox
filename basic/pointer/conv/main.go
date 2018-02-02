package main

import (
	"fmt"
	"time"
	"unsafe"
)

type entryHdr struct {
	accessTime uint32
	keyLen     uint16
	deleted    bool
	testId     uint8
	data       []byte
}

const ENTRY_HDR_SIZE = 8

func main() {

	//测试数据
	testBytes := []byte("It's Friday!")

	//计算结构的总长度
	length := ENTRY_HDR_SIZE + len(testBytes)

	//测试slice, 也可以使用数组，前提是struct的属性只包含基本类型
	var hdrBuf []byte = make([]byte, length)
	hdr := (*entryHdr)(unsafe.Pointer(&hdrBuf[0]))

	//测试赋值
	now := uint32(time.Now().Unix())
	hdr.accessTime = now
	hdr.testId = 12
	hdr.keyLen = 20
	hdr.deleted = true
	hdr.data = testBytes

	//测试转换输出
	newHDR := (*entryHdr)(unsafe.Pointer(&hdrBuf[0]))
	fmt.Printf("------ > %v\n", newHDR.accessTime)
	fmt.Printf("------ > %v\n", newHDR.testId)
	fmt.Printf("------ > %v\n", newHDR.keyLen)
	fmt.Printf("------ > %v\n", newHDR.deleted)
	fmt.Printf("------ > %s\n", newHDR.data)
}
