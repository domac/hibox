package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

const debug = false

//获取普通对象的类型
func getObjectType(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj)
}

//获取普通对象的值
func getObjectValue(obj interface{}) reflect.Value {
	return reflect.ValueOf(obj)
}

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}
func f(target io.Writer) {
	if !getObjectValue(target).IsNil() && getObjectType(target).Kind() == reflect.Ptr {
		fmt.Println("surprise!")
	}
}
