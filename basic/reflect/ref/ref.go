package ref

import (
	//"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	Age  int
}

var handler = func(u *User, message string) {
	//fmt.Printf("Hello, My name is %s, I am %d years old ! so, %s\n", u.Name, u.Age, message)
}

func FiltName(u *User, message string) {
	fn := reflect.ValueOf(handler)
	uv := reflect.ValueOf(u)
	name := uv.Elem().FieldByName("Name")
	name.SetString("XXX")
	fn.Call([]reflect.Value{uv, reflect.ValueOf(message)})
}

var offset uintptr = 0xFFFF

func FiltNameWithReuseOffset(u *User, message string) {
	if offset == 0xFFFF {
		t := reflect.TypeOf(u).Elem()
		name, _ := t.FieldByName("Name")
		offset = name.Offset

	}
	p := (*[2]uintptr)(unsafe.Pointer(&u))
	px := (*string)(unsafe.Pointer(p[0] + offset))
	*px = "YYY"
	fn := reflect.ValueOf(handler)
	uv := reflect.ValueOf(u)
	fn.Call([]reflect.Value{uv, reflect.ValueOf(message)})
}

var cache = map[*uintptr]map[string]uintptr{}

func FiltNameWithCache(u *User, message string) {
	itab := *(**uintptr)(unsafe.Pointer(&u))

	m, ok := cache[itab]
	if !ok {
		m = make(map[string]uintptr)
		cache[itab] = m
	}

	offset, ok := m["Name"]
	if !ok {
		t := reflect.TypeOf(u).Elem()
		name, _ := t.FieldByName("Name")
		offset = name.Offset
		m["Name"] = offset
	}
	p := (*[2]uintptr)(unsafe.Pointer(&u))
	px := (*string)(unsafe.Pointer(p[0] + offset))
	*px = "ZZZ"
	fn := reflect.ValueOf(handler)
	uv := reflect.ValueOf(u)
	fn.Call([]reflect.Value{uv, reflect.ValueOf(message)})

}
