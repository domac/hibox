package brand

//位向量的实现方法大体是：多个位组成一个基本数据类型，基本数据类型组合成数组

const bitsPerWord uint32 = 32 //使用的基本数据类型为32位，int类型
const shift uint32 = 5        //与确定位处于哪个数组元素有关 : 2^5=32,表示移位
const mask uint32 = 0x1F      //与确定位处于数组元素哪一位有关: 二进制 11111
const n uint32 = 30000000     //位长度：向量元素的个数

type Vector struct {
	a [1 + (n-1)/bitsPerWord]uint32
}

func NewVector() *Vector {
	return &Vector{
		a: [1 + (n-1)/bitsPerWord]uint32{},
	}
}

func (self *Vector) set(i uint32) {
	self.a[i>>shift] |= (1 << (i & mask))
}

func (self *Vector) clr(i uint32) {
	self.a[i>>shift] &= ^(1 << (i & mask))
}

func (self *Vector) test(i uint32) bool {
	return (self.a[i>>shift] & (1 << (i & mask))) > 0
}
