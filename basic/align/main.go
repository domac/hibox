package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	b byte
	a int32
	x int64
}
type Data1 struct {
	b byte
	x int64
	a int32
}

func main() {
	var d Data
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align())

	var d1 Data
	t1 := reflect.TypeOf(d1)
	fmt.Println(t1.Size(), t1.Align())

	//f ,_ := t.FieldByName("b")
	//fmt.Println(f.Type.FieldAlign())
}
