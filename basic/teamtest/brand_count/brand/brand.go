package brand

import (
	"bufio"
	"fmt"
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
	for s.Scan() {
		b := s.Bytes()
		lenB := len(b)
		index1 := lasIndexN(b, 9, 32)
		//onlineDate := b[index1+1:]

		if b[index1+3] != 49 {
			continue
		}

		age := b[index1+4]

		if age != 54 && age != 53 && age != 52 && age != 51 && age != 50 && age != 49 {
			continue
		}

		len1 := (lenB - index1 - 1)
		index2 := lasIndexN(b, len1+2, 32)
		price := b[index2+1 : index1]

		len2 := index1 - index2 - 1
		len21 := len1 + len2
		index3 := lasIndexN(b, len21+6, 32)
		//store := b[index3+1 : index2]

		//not equals 'N'
		if b[index3+5] != 78 {
			continue
		}

		len3 := index2 - index3 - 1
		index4 := lasIndexN(b, len21+len3+5, 32)
		name := b[:index4]

		hashKey := hashBytes(name)
		if xh, ok := BRANDKEYS[hashKey]; ok {
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
			compareTopList()
		}
	}
}

func compareTopList() {
	minItem := toplist[0]
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
	toplist[0], toplist[minidx] = toplist[minidx], toplist[0]
}

//输出结果
func ListResult() {
	fmt.Println("------- finish -------")
	values := []BrandItem{}
	for _, item := range toplist {
		if item.xh < 0 {
			continue
		}
		item.TotalValue = BRANDDB[item.xh]
		values = append(values, item)
	}
	quickSort(values, 0, len(values)-1)
	compareSort(values)
	for i, item := range values {
		if item.xh < 0 {
			continue
		}
		fmt.Printf("(%d) name: %s | value: %d | xh: %d\n", (i + 1), item.Name, item.TotalValue, item.xh)
	}
}

func compareSort(arr []BrandItem) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].TotalValue == arr[i+1].TotalValue {
			if arr[i].xh > arr[i+1].xh {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
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
