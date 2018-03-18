package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//测试包
//在数据包中添加长度字段
type TestPacket struct {
	Seq    [2]byte
	Length int16

	//数据包业务数据
	Timestamp      int64
	HostnameLength int16
	Hostname       []byte
	TagLength      int16
	Tag            []byte
	Data           []byte
}

func (p *TestPacket) Pack(writer io.Writer) (err error) {
	err = binary.Write(writer, binary.BigEndian, &p.Seq)
	err = binary.Write(writer, binary.BigEndian, &p.Length)
	err = binary.Write(writer, binary.BigEndian, &p.Timestamp)
	err = binary.Write(writer, binary.BigEndian, &p.HostnameLength)
	err = binary.Write(writer, binary.BigEndian, &p.Hostname)
	err = binary.Write(writer, binary.BigEndian, &p.TagLength)
	err = binary.Write(writer, binary.BigEndian, &p.Tag)
	err = binary.Write(writer, binary.BigEndian, &p.Data)
	return err
}

func (p *TestPacket) Unpack(reader io.Reader) (err error) {
	err = binary.Read(reader, binary.BigEndian, &p.Seq)
	err = binary.Read(reader, binary.BigEndian, &p.Length)
	err = binary.Read(reader, binary.BigEndian, &p.Timestamp)
	err = binary.Read(reader, binary.BigEndian, &p.HostnameLength)
	p.Hostname = make([]byte, p.HostnameLength)
	err = binary.Read(reader, binary.BigEndian, &p.Hostname)
	err = binary.Read(reader, binary.BigEndian, &p.TagLength)
	p.Tag = make([]byte, p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Tag)
	p.Data = make([]byte, p.Length-8-2-p.HostnameLength-2-p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Data)
	return err
}

func (p *TestPacket) String() string {
	return fmt.Sprintf("seq:%s length:%d timestamp:%d hostname:%s tag:%s data:%s",
		p.Seq,
		p.Length,
		p.Timestamp,
		p.Hostname,
		p.Tag,
		p.Data,
	)
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	testData := fmt.Sprintf("time now is %s", time.Now().Format("2006-01-02 15:04:05"))

	packet := &TestPacket{
		Seq:            [2]byte{'S', '1'},
		Timestamp:      time.Now().Unix(),
		HostnameLength: int16(len(hostname)),
		Hostname:       []byte(hostname),
		TagLength:      4,
		Tag:            []byte("Test"),
		Data:           []byte(testData),
	}

	packet.Length = 8 + 2 + packet.HostnameLength + 2 + packet.TagLength + int16(len(packet.Data))

	buf := new(bytes.Buffer)
	packet.Pack(buf)
	packet.Pack(buf)
	packet.Pack(buf)
	packet.Pack(buf)
	packet.Pack(buf)

	scanner := bufio.NewScanner(buf)

	//自定义分隔函数
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'S' {
			if len(data) > 4 {
				length := int16(0)
				binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
				if int(length)+4 <= len(data) {
					bodyLen := int(length) + 4
					return bodyLen, data[:bodyLen], nil
				}
			}
		}
		return
	})

	//扫描遍历
	for scanner.Scan() {
		sp := new(TestPacket)
		sp.Unpack(bytes.NewReader(scanner.Bytes()))
		log.Println(sp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("无效数据包")
	}
}
