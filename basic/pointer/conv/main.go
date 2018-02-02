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
}

const ENTRY_HDR_SIZE = 8

func main() {
	var hdrBuf [ENTRY_HDR_SIZE]byte
	hdr := (*entryHdr)(unsafe.Pointer(&hdrBuf[0]))

	now := uint32(time.Now().Unix())
	hdr.accessTime = now
	hdr.testId = 12
	hdr.keyLen = 20
	hdr.deleted = true

	newHDR := (*entryHdr)(unsafe.Pointer(&hdrBuf[0]))
	fmt.Printf("------ > %v\n", newHDR)
}
