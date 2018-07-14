package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//netstat -an | grep 37070 | awk  '/TIME_WAIT/ {print $6}' | sort | uniq -c

const URL = "http://localhost:37070/hello"

func PrintLocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}

	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}

func doGet(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%d: %s -- %v\n", id, string(buf), err)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: PrintLocalDial,
			//MaxIdleConns: 128,
			MaxIdleConnsPerHost: 128,
		},
	}
	for {
		for i := 0; i < 100; i++ {
			go doGet(client, URL, i)
		}
		time.Sleep(2 * time.Second)
	}
}
