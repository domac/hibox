package main

import (
	"fmt"
	"runtime"
)

func Foo() {
	fmt.Println("call foo:", printMyName(1), "|", printMyName(2))
	Boo()
}

func Boo() {
	fmt.Println("call boo:", printMyName(1), "|", printMyName(2))
}

func printMyName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name()
}

func main() {
	Foo()
}
