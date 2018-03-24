package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type ReqMsg struct {
	Code int
	Req  []byte
}

type RespMsg struct {
	Code int
	Resp []byte
}

type SendStruct struct{}

func (s *SendStruct) OnSend(req *ReqMsg, resp *RespMsg) error {
	code := req.Code + 1000
	ret := string(req.Req) + "_vertify"

	resp.Code = code
	resp.Resp = []byte(ret)

	return nil
}

func main() {
	req := new(SendStruct)
	rpc.Register(req)

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err.Error())
	}
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
