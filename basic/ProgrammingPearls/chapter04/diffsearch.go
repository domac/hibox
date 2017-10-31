package main

import ()
import "fmt"

//进行二分搜索的数组要求是有序的，二分搜索通过持续跟踪数组中包含元素t的范围来解决问题。
// 一开始，这个范围是整个数组，然后通过将t与数组的中间项进行比较并抛弃一半的范围来缩小范围。
//该过程持续进行，直到在数组中找到t或确定包含t的范围为空时为止，在有n个元素的表中，二分搜索大约需要执行log2n次比较操作
func binSearch(srcArray []int, des int) (int, int) {
	low := 0
	count := 0
	high := len(srcArray) - 1
	for low <= high {
		count++
		middle := low + ((high - low) >> 1) //这里是为了溢出
		if des == srcArray[middle] {
			return middle, count
		} else if des < srcArray[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -(low + 1), count
}

//下面需要让二分搜索更快
//差值搜索
func diffSearch(srcArray []int, des int) (int, int) {
	count := 0
	low := 0
	high := len(srcArray) - 1
	for low <= high {

		//首尾相同的情况
		if srcArray[low] == srcArray[high] {
			count++
			if srcArray[low] == des {
				return low, count
			} else {
				return -1, count
			}
		}

		middle := low + (high-low)*(des-srcArray[low])/(srcArray[high]-srcArray[low])
		count++
		if des < srcArray[middle] {
			high = middle - 1
		} else if des > srcArray[middle] {
			low = middle + 1
		} else {
			return middle, count
		}
	}

	return -(low + 1), count
}

func main() {
	testArray := []int{1, 3, 5, 6, 7, 9, 13, 14, 21, 22, 24, 37, 38, 39, 40, 51, 58, 61, 68, 77, 79, 80, 83, 91, 98}
	id1, c1 := binSearch(testArray, 14)
	fmt.Printf("二分搜索结果:%d  比较次数：%d\n", id1, c1)

	id2, c2 := diffSearch(testArray, 14)
	fmt.Printf("二分搜索结果(改进):%d  比较次数：%d\n", id2, c2)
}
