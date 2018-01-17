//+build debug

package main

import (
	"fmt"
)

const Dev = true

func Debug(a ...interface{}) {
	fmt.Println(a...)
}
