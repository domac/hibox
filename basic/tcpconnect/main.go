package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4242")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go copyToStderr(conn)
	}
}

func copyToStderr(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Copied %d bytes; finish with err = %v", n, err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
}
