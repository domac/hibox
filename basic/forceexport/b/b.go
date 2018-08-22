package b

import (
	_ "github.com/domac/hibox/basic/forceexport/a"
	_ "unsafe"
)

//go:linkname privateSay a.say
func privateSay(name string) string

func PublicSay(name string) string {
	return privateSay(name)
}
