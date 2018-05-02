package main

import (
	"fmt"
	scanner "github.com/domac/hibox/basic/sqlaw/scan"
	"github.com/zserge/webview"
	"log"
	"net"
	"net/http"
	"strings"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var indexHTML = `
<!doctype html>
<html>
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	</head>
	<body>
		<br>
		<button onclick="external.invoke('opendir')">选择目录</button>
		<input id="scandir" type="text" />
		<button onclick="external.invoke('changeDir:'+document.getElementById('scandir').value)">
			扫描
		</button>
		<!-- <input id="new-color" value="#e91e63" type="color" /> -->
		<div id='checkresult' style="width: 100%;height: 530px;overflow-y: auto;margin-top: 10px;background-color: black;color: #33FF00">
		</div>
	</body>
</html>
`

func startServer() string {

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(indexHTML))
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

//终端打印
func consolePrint(w webview.WebView, format string, args ...interface{}) {
	output := format
	if len(args) > 0 {
		output = fmt.Sprintf(format, args)
	}

	w.Eval("document.getElementById('checkresult').innerText='" + output + "'")
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case data == "opendir":
		dir := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		log.Println("open dir = ", dir)
		w.Eval("document.getElementById('scandir').value='" + dir + "'")
	case strings.HasPrefix(data, "changeDir:"):
		scanDir := strings.TrimPrefix(data, "changeDir:")
		//consolePrint(w, scanDir)
		res := scanner.DoScan([]string{scanDir})
		//println(string(res))
		//consolePrint(w, "完成:"+string(res))
		for _, s := range res {
			s += "<br>"
			w.Eval("document.getElementById('checkresult').innerHTML+='" + s + "'")
		}
		//consolePrint(w, out)

	}
}

func main() {
	url := startServer()
	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "SQL注入分析工具",
		Resizable: true,
		URL:       url,
		ExternalInvokeCallback: handleRPC,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()
}
