package main

import (
	"fmt"
	"runtime"
)

func CallerName(skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

func main() {

	for skip := 0; ; skip++ {
		name, file, line, ok := CallerName(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v\n", skip)
		fmt.Printf("  file = %v, line = %d\n", file, line)
		fmt.Printf("  name = %v\n", name)
	}

}
