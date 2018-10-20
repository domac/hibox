package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func main() {

	var buf [md5.Size + 4]byte
	key := []byte("test key")

	rand.Read(buf[:8])

	fmt.Printf("%v\n", buf[:8])

	hash := md5.New()

	hash.Write(buf[:8])
	hash.Write(key)

	verify := hash.Sum(nil)
	fmt.Printf("verify = %v\n", verify)

	binary.BigEndian.PutUint32(buf[md5.Size:], uint32(123))

	copy(buf[:md5.Size], verify)

	fmt.Printf("buf = %v\n", buf)

	serverid := binary.BigEndian.Uint32(buf[md5.Size:])
	fmt.Printf("serverid = %v\n", serverid)
}
