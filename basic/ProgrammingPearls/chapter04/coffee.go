package main

func main() {
	//初始化
	whiteBeans := 101
	totalBeans := 2*whiteBeans + 1

	blackBeans := 0

	n := whiteBeans % 2

	if n == 1 { // w = 2n +1

		blackBeans = totalBeans - (2*n + 1)

	} else { // w = 2n

		blackBeans = totalBeans - 2*n
	}

	if n == 1 { //有白

		if blackBeans%2 == 0 {
			println("最后一个为黑色")
		} else {
			println("最后一个为白色")
		}

	} else {
		println("最后一个为黑色") //都是黑
	}
}
