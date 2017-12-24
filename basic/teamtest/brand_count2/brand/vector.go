package brand

//位向量的实现方法大体是：多个位组成一个基本数据类型，基本数据类型组合成数组

const bitsPerWord uint32 = 32 //使用的基本数据类型为32位，int类型
const shift uint32 = 5        //与确定位处于哪个数组元素有关 : 2^5=32,表示移位
const mask uint32 = 0x1F      //与确定位处于数组元素哪一位有关: 二进制 11111
const n uint32 = 365 * 12     //位长度：向量元素的个数

type Vector struct {
	arr   [1 + (n-1)/bitsPerWord]uint32
	count int
}

func NewVector() *Vector {
	return &Vector{}
}

//设值
//因为是置1操作，所以只需要把原数组元素与位位置掩码进行或操作即可
func (self *Vector) Set(i uint32) {
	self.arr[i>>shift] |= (1 << (i & mask))
}

//清零
//因为是置0操作，所以先把位位置掩码取非，使特定位位0，再与原数组元素进行与操作即可
func (self *Vector) Clr(i uint32) {
	self.arr[i>>shift] &= ^(1 << (i & mask))
}

//判断特定位，只需要原始数组元素与位位置掩码进行与操作，即可判断当前位是否为1
func (self *Vector) Test(i uint32) bool {
	return (self.arr[i>>shift] & (1 << (i & mask))) > 0
}
