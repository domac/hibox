package main

const bitsPerWord uint32 = 32 //int 有32位
const shift uint32 = 5        //2^5=32,表示移位
const mask uint32 = 0x1F      //二进制 1111
const n uint32 = 10000000     //向量元素的个数

//向量数组i
var a = [1 + n/bitsPerWord]uint32{}

//设值
func set(i uint32) {
	a[i>>shift] |= (1 << (i & mask))
}

//清零
func clr(i uint32) {
	a[i>>shift] &= ^(1 << (i & mask))
}

//测试输出
func test(i uint32) uint32 {
	return a[i>>shift] & (1 << (i & mask))
}

//下面的版本实现效果是一样的

//设值
func set0(i uint32) {
	a[i/32] |= (1 << (i % 32))
}

//清零
func clr0(i uint32) {
	a[i/32] &= ^(1 << (i % 32))
}

//测试输出
func test0(i uint32) uint32 {
	return a[i/32] & (1 << (i % 32))
}

func main() {

	ti := uint32(1026)
	println(test(ti))
	println("------------")
	set(ti)
	println(test(ti))
	clr(ti)
	println(test(ti))

}
