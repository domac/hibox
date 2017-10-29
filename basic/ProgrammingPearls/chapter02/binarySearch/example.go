package main

/**
 * 参考资料:
 * http://www.jianshu.com/p/684e286fa856
 * https://segmentfault.com/q/1010000000359749
 * http://www.cnblogs.com/wuyuegb2312/archive/2013/05/26/3090369.html
 */

//二分查找法的O(log n)让它成为十分高效的算法。不过它的缺陷就在它的限定之上：
//	1、必须有序，我们很难保证我们的数组都是有序的。当然可以在构建数组的时候进行排序，可是又落到了第二个瓶颈上：它必须是数组。
//	2、数组读取效率是O(1)，可是它的插入和删除某个元素的效率却是O(n)。因而导致构建有序数组变成低效的事情。
//解决这些缺陷问题更好的方法应该是使用二叉查找树; 最好自然是自平衡二叉查找树了，自能高效的（O(n log n)）构建有序元素集合，又能如同二分查找法一样快速（O(log n)）的搜寻目标数

/*
 * 一、非递归方式完成二分查找法
 * 参数:整型数组,需要比较的数.
 * 算法的思想就是：从数组中间开始，每次排除一半的数据，时间复杂度为O(lgN)。这依赖于数组有序这个性质。如果t存在数组中，则返回t在数组的位置；否则，不存在则返回-(l+1)。
 * 这里需要解释下为什么t不存在数组中时不是返回-1而要返回-(l+1)。首先我们可以观察l的值，如果查找不成功，则l的值恰好是t应该在数组中插入的位置。
 */
func binSearch1(srcArray []int, des int) int {
	//第一个位置.
	low := 0
	//最高位置.数组最大坐标-1
	high := len(srcArray) - 1
	for low <= high {
		//中间位置计算,low+ 最高位置减去最低位置,右移一位,相当于除2.也可以用(high+low)/2
		middle := low + ((high - low) >> 1) //这里是为了溢出
		//与最中间的数字进行判断,是否相等,相等的话就返回对应的数组下标.
		if des == srcArray[middle] {
			return middle
			//如果小于的话则移动最高层的"指针"
		} else if des < srcArray[middle] {
			high = middle - 1
			//移动最低的"指针"
		} else {
			low = middle + 1
		}
	}
	return -(low + 1)
}

/**
 * 二、递归方法实现二分查找法.
 * 参数: srcArray数组
 * 参数: low 数组第一位置
 * 参数: high 最高
 * 参数: des 要查找的值.
 */
func binSearch2(srcArray []int, low int, high int, des int) int {
	if low <= high {
		mid := low + (high-low)>>1 //这里是为了溢出; 也可以用(high+low)/2,
		if des == srcArray[mid] {
			return mid
		} else if des < srcArray[mid] {
			//移动low和high
			return binSearch2(srcArray, low, mid-1, des)
		} else if des > srcArray[mid] {
			return binSearch2(srcArray, mid+1, high, des)
		}
	}
	return -1
}

func main() {
	testArray := []int{1, 3, 5, 6, 8, 11, 13, 18, 21, 22, 24, 37}

	//测试方法1
	i1 := binSearch1(testArray, 21)
	println(i1)

	//测试方法2
	i2 := binSearch2(testArray, 0, len(testArray), 21)
	println(i2)
}
