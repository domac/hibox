// +build go1.10

package gid

import "unsafe"

type g struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	_panic       uintptr
	_defer       uintptr
	m            uintptr
	sched        gobuf
	syscallsp    uintptr
	syscallpc    uintptr
	stktopsp     uintptr
	param        unsafe.Pointer
	atomicstatus uint32
	stackLock    uint32
	goid         int64
}

type stack struct {
	lo uintptr
	hi uintptr
}

type gobuf struct {
	fields [7]uintptr
}
