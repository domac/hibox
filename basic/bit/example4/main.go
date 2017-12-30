package main

import (
	"fmt"
	"strings"
)

const (
	UPPER = 1
	LOWER = 2
	CAP   = 4
	REV   = 8
)

func main() {
	fmt.Println(procstr("HELLO LIHAOQUAN!", LOWER|REV|CAP))
}

func procstr(str string, conf byte) string {

	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)

		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}

	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}

	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}

	if (conf & REV) != 0 {
		str = rev(str)
	}

	return str
}
