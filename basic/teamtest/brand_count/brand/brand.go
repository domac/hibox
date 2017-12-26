package brand

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//全局常量
const TOPNUM = 40

//全局变量
var (
	BRANDKEYS = make(map[uint64]int, 30000000)
	BRANDDB   = []int{}
	toplist   [TOPNUM]BrandItem
	topMap    = make(map[uint64]int)
)

type BrandItem struct {
	Name    string
	HashKey uint64
	xh      int

	TotalValue int
}

func InitKeys(dataFile string) error {
	f, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	idx := 0
	for s.Scan() {
		b := s.Bytes()
		//逆向切割
		hashKey := hashBytes(b)
		BRANDKEYS[hashKey] = idx
		idx++
	}
	keysLen := idx
	BRANDDB = make([]int, keysLen, keysLen)
	for i := 0; i < keysLen; i++ {
		BRANDDB[i] = 0
	}

	for i := 0; i < TOPNUM; i++ {
		toplist[i] = BrandItem{
			xh: -1,
		}
	}
	return nil
}

//数据文件读入处理
func ReadAndHandle(dataFile string) error {
	fmt.Println("------- start -------")
	f, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(MyScanLines)
	for s.Scan() {
		b := s.Bytes()
		index1 := lasIndexN(b, 9, 32)
		//onlineDate := b[index1+1:]

		// if b[index1+3] != 49 {
		// 	continue
		// }

		age := b[index1+4]

		//11~16
		if age != 54 && age != 53 && age != 52 && age != 51 && age != 50 && age != 49 {
			continue
		}

		index2 := lasIndexIdx(b, index1, 32)
		index3 := lasIndexIdx(b, index2, 32)

		//not equals 'N'
		if b[index3+5] != 78 {
			continue
		}

		index4 := lasIndexIdx(b, index3, 32)
		name := b[:index4]

		hashKey := hashBytes(name)
		if xh, ok := BRANDKEYS[hashKey]; ok {
			price := b[index2+1 : index1]
			currentValue := BRANDDB[xh] + parsebyteToInt(price)
			BRANDDB[xh] = currentValue
			updateTopList(name, hashKey, xh, currentValue)
		}
	}
	return nil
}

func updateTopList(name []byte, hashKey uint64, xh, currentValue int) {

	flag, ok := topMap[hashKey]
	if !ok || flag == 0 {

		if len(topMap) < TOPNUM {
			compareTopList()
		}

		minItem := toplist[0]

		isReplace := false

		minItemTotalValue := 0
		if minItem.xh >= 0 {
			minItemTotalValue = BRANDDB[minItem.xh]
		}

		if minItemTotalValue < currentValue {
			isReplace = true
		} else if minItemTotalValue == currentValue {
			if minItem.xh > xh {
				isReplace = true
			}
		}
		if isReplace {
			tempKey := minItem.HashKey
			minItem.Name = string(name)
			minItem.HashKey = hashKey
			minItem.xh = xh
			toplist[0] = minItem
			topMap[tempKey] = 0
			topMap[hashKey] = 1

		}
	} else {
		compareTopList()
	}
}

func compareTopList() {
	minItem := toplist[0]

	if minItem.xh < 0 {
		return
	}

	minidx := 0
	ilen := len(toplist)
	for i := 1; i < ilen; i++ {
		temp := toplist[i]
		tempVal := 0
		if temp.xh >= 0 {
			tempVal = BRANDDB[temp.xh]
		}

		minItemTotalValue := 0
		if minItem.xh >= 0 {
			minItemTotalValue = BRANDDB[minItem.xh]
		}

		if tempVal < minItemTotalValue {
			minItem = temp
			minidx = i
		} else if tempVal == minItemTotalValue {
			if temp.xh > minItem.xh {
				minItem = temp
				minidx = i
			}
		}

	}

	if minidx > 0 {
		toplist[0], toplist[minidx] = toplist[minidx], toplist[0]
	}
}

//输出结果
func ListResult() {

	values := []BrandItem{}
	for _, item := range toplist {
		if item.xh < 0 {
			continue
		}
		item.TotalValue = BRANDDB[item.xh]
		values = append(values, item)
	}

	if len(values) == 0 {
		log.Fatal("result is null \n")
		return
	}

	log.Println(">> quick sort")
	quickSort(values, 0, len(values)-1)

	log.Println(">> xh sort")
	values = compareSort(values)
	for i, item := range values {
		if item.xh < 0 {
			continue
		}
		fmt.Printf("(%d) name: %s | value: %d | xh: %d\n", (i + 1), item.Name, item.TotalValue, item.xh)
	}
	fmt.Println("------- finish -------")
}

func compareSort(arr []BrandItem) []BrandItem {
	lenS := len(arr)
	currentTotalValue := arr[0].TotalValue
	currentStartIndex := 0
	for i := 1; i < lenS; i++ {
		targetTotalValue := arr[i].TotalValue
		if currentTotalValue > targetTotalValue || i == lenS-1 {
			currentTotalValue = targetTotalValue
			quickSubXhArray(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func quickSubXhArray(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].xh
		for i <= j {
			for arr[i].xh < key {
				i++
			}
			for arr[j].xh > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSubXhArray(arr, start, j)
		}
		if end > i {
			quickSubXhArray(arr, i, end)
		}
	}
}

//快速排序:从大到小
func quickSort(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].TotalValue
		for i <= j {
			for arr[i].TotalValue > key {
				i++
			}
			for arr[j].TotalValue < key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}

}
