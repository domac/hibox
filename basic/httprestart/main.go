package main

import (
	"flag"
	"fmt"
	r "github.com/domac/hibox/basic/httprestart/restart"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "Address to listen on!")
var sname = flag.String("name", "master", "node name")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from node [%s]  %v \n", *sname, os.Getpid())
}

func main() {
	flag.Parse()
	fmt.Printf("web server run with addr : %s | %s\n", *sname, *addr)

	//创建或复用listener属性
	ln, err := r.ImportOrCreateListener(*addr)
	if err != nil {
		fmt.Printf("unable to import or create a listener : %v \n", err)
		os.Exit(1)
	}

	http.HandleFunc("/hello", handler)

	//开发web 服务
	server, err := r.StartServer(*addr, ln)
	if err != nil {
		fmt.Printf("Exiting : %v\n", err)
		return
	}

	//事件监听
	err = r.WaitForSignals(server)
	if err != nil {
		fmt.Printf("Exiting : %v\n", err)
		return
	}
	fmt.Println("Existing ")
}
