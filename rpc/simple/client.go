package main

import (
	"fmt"
	"net"
	"net/rpc"
	"runtime"
	"time"
	//"os"
)

type ReqMsg struct {
	Code int
	Req  []byte
}

type RespMsg struct {
	Code int
	Resp []byte
}

func RpcDialWithTimeout(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return rpc.NewClient(conn), nil
}

func main() {
	req := &ReqMsg{1, []byte("hello")}

	for i := 0; i < 20000; i++ {

		//cli, err := rpc.Dial("tcp", "127.0.0.1:8000")

		cli, err := RpcDialWithTimeout("tcp", "127.0.0.1:8000", 10*time.Second)

		gonum := runtime.NumGoroutine()

		if err != nil {
			fmt.Println(err, i, gonum)
			continue
		}

		var resp RespMsg
		cli.Call("SendStruct.OnSend", req, &resp)

		fmt.Printf(">>> code:%d, data :%s", resp.Code, string(resp.Resp))
	}

}
