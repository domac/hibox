package main

func AbsFunc(n int64) int64 {
	y := n >> 63
	return (y ^ n) - y
}

func main() {
	test := int64(-1233)
	println(AbsFunc(test))
}
