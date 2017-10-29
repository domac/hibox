package main

/**
 * 参考资料:
 * http://www.jianshu.com/p/684e286fa856
 * http://www.cnblogs.com/wuyuegb2312/p/3139926.html
 */

//循环不变式主要用来帮助理解算法的正确性。形式上很类似与数学归纳法，它是一个需要保证正确断言。
//对于循环不变式，必须证明它的三个性质：
//	初始化：它在循环的第一轮迭代开始之前，应该是正确的。
//	保持：如果在循环的某一次迭代开始之前它是正确的，那么，在下一次迭代开始之前，它也应该保持正确。
//	终止：循环能够终止，并且可以得到期望的结果。

//二分查找数字第一次出现的位置
//算法分析：设定两个不存在的元素a[-1]和a[n]，使得a[-1] < t <= a[n]
//但是我们并不会去访问者两个元素，因为(l+u)/2 > l=-1, (l+u)/2 < u=n。
//循环不变式为l<u && t>a[l] && t<=a[u] 。
//循环退出时必然有l+1=u, 而且a[l] < t <= a[u]。
//循环退出后u的值为t可能出现的位置，其范围为[0, n]，如果t在数组中，则第一个出现的位置p=u，如果不在，则设置p=-1返回。
//该算法的效率虽然解决了更为复杂的问题，但是其效率比初始版本的二分查找还要高，因为它在每次循环中只需要比较一次，前一程序则通常需要比较两次。
func binSearchFirst(srcArray []int, n, des int) int {
	low := -1
	high := n

	for low+1 != high {
		middle := low + ((high - low) >> 1) //这里是为了溢出

		if des > srcArray[middle] {
			low = middle
		} else {
			high = middle
		}
	}

	/*assert: l+1=u && a[l]<t<=a[u]*/
	p := high
	if p >= n || srcArray[p] != des {
		p = -1
	}
	return p
}

//如果要查找数字在数组中最后出现的位置呢？
func binSearchLast(srcArray []int, n, des int) int {
	low := -1
	high := n

	for low+1 != high {
		middle := low + ((high - low) >> 1) //这里是为了溢出

		if des >= srcArray[middle] {
			low = middle
		} else {
			high = middle
		}
	}

	/*assert: l+1=u && a[l]<t<=a[u]*/
	p := low
	if p <= 1 || srcArray[p] != des {
		p = -1
	}
	return p
}

func main() {
	testArray := []int{1, 3, 5, 6, 7, 7, 13, 13, 21, 22, 24, 37}
	idx := binSearchFirst(testArray, len(testArray), 13)
	println(idx)

	idx2 := binSearchLast(testArray, len(testArray), 13)
	println(idx2)
}
