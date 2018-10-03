package main

import (
	"encoding/binary"
	"fmt"
	//"io"
)

type ByteBuffer struct {
	B []byte
}

func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		//B: make([]byte, 0, 10),
		B: make([]byte, 0, 0),
	}
}

func (b *ByteBuffer) Len() int {
	return len(b.B)
}

// func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
// 	p := b.B
// 	nStart := int64(len(p))
// 	nMax := int64(cap(p))
// 	n := nStart
// 	if nMax == 0 {
// 		nMax = 64
// 		p = make([]byte, nMax)
// 	} else {
// 		p = p[:nMax]
// 	}
// 	for {
// 		if n == nMax {
// 			nMax *= 2
// 			bNew := make([]byte, nMax)
// 			copy(bNew, p)
// 			p = bNew
// 		}
// 		nn, err := r.Read(p[n:])
// 		n += int64(nn)
// 		if err != nil {
// 			b.B = p[:n]
// 			n -= nStart
// 			if err == io.EOF {
// 				return n, nil
// 			}
// 			return n, err
// 		}
// 	}
// }

func (b *ByteBuffer) Bytes() []byte {
	return b.B
}

// Reset makes ByteBuffer.B empty.
func (b *ByteBuffer) Reset() {
	b.B = b.B[:0]
}

//------ 写方法

func (b *ByteBuffer) Write(p []byte) (int, error) {
	b.B = append(b.B, p...)
	return len(p), nil
}

func (b *ByteBuffer) WriteByte(c byte) error {
	b.B = append(b.B, c)
	return nil
}

//------- 读方法
func (b *ByteBuffer) Read(p []byte) (int, error) {
	return copy(p, b.B), nil
}

func main() {
	bb := NewByteBuffer()
	binary.Write(bb, binary.BigEndian, uint32(0))
	fmt.Printf("[debug]: %v\n", bb.B)

	bb.WriteByte('1')

	fmt.Printf("[debug]: %v\n", bb.B)

	bb.WriteByte(byte(128))

	fmt.Printf("[debug]: %v\n", bb.B)

	bb.Write([]byte{'1', '2', '3', '4'})

	fmt.Printf("-[debug]: %v\n", bb.B)

	prefixLen := bb.Len()

	println("prefixLen =", prefixLen)

	//更新size
	binary.BigEndian.PutUint32(bb.B, uint32(bb.Len()))

	fmt.Printf("--[debug]: %v\n", bb.B)

	println("---------------------------")
	binary.Write(bb, binary.BigEndian, uint64(64))
	fmt.Printf("[debug]: %v\n", bb)
	bb.WriteByte('p')
	fmt.Printf("[debug]: %v\n", bb.B)

	println("---------------------------")
	var size uint32
	binary.Read(bb, binary.BigEndian, &size)
	fmt.Printf("[debug]: size=%v\n", size)
	fmt.Printf("[debug]: data=%v\n", bb.B)

	println("---------------------------")

	prefixLen = bb.Len()

	println("current prefixLen =", prefixLen)

	//更新size
	binary.BigEndian.PutUint32(bb.B, uint32(bb.Len()))

	fmt.Printf("[debug]: %v\n", bb.B)

	binary.BigEndian.PutUint16(bb.B[4:], uint16(8))

	fmt.Printf("[debug]: %v\n", bb.B)

}
