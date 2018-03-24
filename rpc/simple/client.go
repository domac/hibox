package main

import (
	"fmt"
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

func main() {
	req := &ReqMsg{1, []byte("hello")}

	cli, err := rpc.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var resp RespMsg
	cli.Call("SendStruct.OnSend", req, &resp)

	fmt.Printf("code:%d, data :%s", resp.Code, string(resp.Resp))
}
