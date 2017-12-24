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
	BRANDKEYS  = make(map[uint64]uint32, 30000000)
	ONLINESMAP = make(map[uint64]uint32, 70000000)
	BRANDDB    = []int{}

	dataList  = []uint32{}
	namedList = []string{}
)

type BrandItem struct {
	Name    string
	HashKey uint64
	xh      uint32

	DateCount  uint32
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
		if b := s.Bytes(); b != nil {
			//逆向切割
			hashKey := hashBytes(b)
			BRANDKEYS[hashKey] = uint32(idx)
			idx++
		}
	}
	keysLen := idx
	BRANDDB = make([]int, keysLen, keysLen)
	dataList = make([]uint32, keysLen, keysLen)
	namedList = make([]string, keysLen, keysLen)

	for i := 0; i < keysLen; i++ {
		BRANDDB[i] = 0
		dataList[i] = 0
		namedList[i] = ""
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
		index1 := lasIndexN(b, 9, 32)
		index2 := lasIndexIdx(b, index1, 32)
		index3 := lasIndexIdx(b, index2, 32)
		index4 := lasIndexIdx(b, index3, 32)

		//基础数据
		name := b[:index4]
		hashKey := hashBytes(name)
		if xh, ok := BRANDKEYS[hashKey]; ok {
			price := b[index2+1 : index1]
			BRANDDB[xh] += parsebyteToInt(price)
			onlineDate := b[index1+1:]
			combineHashHey := combinehashBytes32(onlineDate, xh)

			item := ONLINESMAP[combineHashHey] + 1
			ONLINESMAP[combineHashHey] = item

			dv := dataList[xh]

			if item > dv {
				dataList[xh] = item
			}

			if dv == 0 {
				namedList[xh] = string(name)
			}

			name = name[:0]
			price = price[:0]
			onlineDate = onlineDate[:0]
			name = nil
			price = nil
			onlineDate = nil
		}
	}
	return nil
}

//输出结果
func ListResult() {
	ONLINESMAP = make(map[uint64]uint32, 0)
	ONLINESMAP = nil

	values := make([]BrandItem, len(BRANDKEYS), len(BRANDKEYS))

	cid := 0
	for _, idx := range BRANDKEYS {
		d := dataList[idx]
		tv := BRANDDB[idx]
		name := namedList[idx]
		if d > 1 {
			values[cid] = BrandItem{Name: name, TotalValue: tv, DateCount: d, xh: idx}
			cid++
		}
	}
	values = values[:cid]
	log.Println(">> quick sort")

	quickSort(values, 0, len(values)-1)

	newCount := cid
	values2 := []BrandItem{}
	for i := 0; i < newCount; i++ {
		vi := values[i]
		values2 = append(values2, vi)
	}

	log.Println(">> value sort")
	values2 = compareSortValue(values2)

	log.Println(">> xh sort")
	values2 = compareSortXh(values2)

	for i := 0; i < TOPNUM; i++ {
		currentItem := values2[i]
		fmt.Printf("%d) %s, dateCount: %d, value: %d, xh: %d \n", i+1, currentItem.Name, currentItem.DateCount, currentItem.TotalValue, currentItem.xh)
	}
	fmt.Println("------- finish -------")
}

func compareSortValue(arr []BrandItem) []BrandItem {
	lenS := len(arr)
	currentDateCount := arr[0].DateCount
	currentStartIndex := 0
	for i := 1; i < lenS; i++ {
		targetDateCount := arr[i].DateCount
		if currentDateCount > targetDateCount || i == lenS-1 {
			currentDateCount = targetDateCount
			quickSubValueArray(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func compareSortXh(arr []BrandItem) []BrandItem {
	lenS := len(arr)
	currentDateCount := arr[0].DateCount
	currentTotalValue := arr[0].TotalValue
	currentStartIndex := 0

	for i := 1; i < lenS; i++ {
		targetDateCount := arr[i].DateCount
		targetTotalValue := arr[i].TotalValue

		if currentDateCount > targetDateCount || currentTotalValue > targetTotalValue || i == lenS-1 {
			currentDateCount = targetDateCount
			currentTotalValue = targetTotalValue
			quickSubXhArray(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func quickSubValueArray(arr []BrandItem, start, end int) {
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
			quickSubValueArray(arr, start, j)
		}
		if end > i {
			quickSubValueArray(arr, i, end)
		}
	}
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

func quickSort(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].DateCount
		for i <= j {
			for arr[i].DateCount > key {
				i++
			}
			for arr[j].DateCount < key {
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
