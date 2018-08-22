package a

import (
	_ "unsafe"
)

//go:linkname say a.say
func say(name string) string {
	return "hello, " + name
}
