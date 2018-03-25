package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const (
	UA  = "hibox"
	URL = "https://pic1.zhimg.com/80/v2-1e6430bbecf17b1a74354746476a15c3_r.jpg"
)

func main() {
	wd, _ := os.Getwd()
	filePath := path.Join(wd, "test.jpg")
	fmt.Printf("file path : %s\n", filePath)

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileStat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	ss := fileStat.Size()
	fmt.Printf("file size : %d\n", ss)

	f.Seek(ss, 0)
	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Set("UserAgent", UA)

	//先测试请求前1000byte
	bs := fmt.Sprintf("bytes=%d-1000", ss)
	req.Header.Set("Range", bs)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	written, err := io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	println("written: ", written)

	//再测试从1000byte后处理剩余的请求
	bs = fmt.Sprintf("bytes=1001-")
	req.Header.Set("Range", bs)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	written, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}
	println("written2: ", written)
	defer resp.Body.Close()
}
