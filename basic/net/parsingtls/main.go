package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//go copyToStderr(conn)
		go logSNI(conn)
	}
}

func logSNI(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, conn, 1+2+2); err != nil {
		log.Println(err)
		return
	}

	length := binary.BigEndian.Uint16(buf.Bytes()[3:5])
	if _, err := io.CopyN(&buf, conn, int64(length)); err != nil {
		log.Println(err)
		return
	}

	log.Printf("Got a connection with SNI %s\n", buf.Bytes())

}

func proxy(conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", "gophercon.com:443")
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	go io.Copy(remote, conn)
	io.Copy(conn, remote)
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
