package main

import (
	"log"
	"strconv"
)

func SetRedisProtocol(buf []byte, args ...string) []byte {
	buf = append(buf, '*')
	buf = strconv.AppendInt(buf, int64(len(args)), 10)
	buf = append(buf, '\r', '\n')
	for _, arg := range args {
		buf = append(buf, '$')
		buf = strconv.AppendInt(buf, int64(len(arg)), 10) //mykey
		buf = append(buf, '\r', '\n')
		buf = append(buf, arg...) //myvalue
		buf = append(buf, '\r', '\n')
	}
	return buf
}

func main() {

	var buf []byte

	buf = SetRedisProtocol(buf, "SET", "testkey", "123")

	log.Printf("->\n%s\n", buf)

	println("pipeline : ")
	buf = buf[:0]

	for i := 0; i < 10; i++ {
		buf = SetRedisProtocol(buf, "SET", "testkey", strconv.Itoa(i))
	}
	log.Printf("->\n%s\n", buf)
}
