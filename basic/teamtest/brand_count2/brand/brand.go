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
	BRANDKEYS    = make(map[uint64]int, 1024*1024*20)
	BRANDDB      = make(map[uint64]uint64, 1024*1024)
	topMap       = make(map[uint64]int)
	toplist      [TOPNUM]BrandItem
	ONLINESMAP   = make(map[uint64]bool)
	ONLINESCOUNT = make(map[uint64]uint64)
)

type BrandItem struct {
	Name        string
	TotalValue  uint64
	HashKey     uint64
	OnlineCount uint64

	seq int
}

func InitKeys(dataFile string) error {

	f, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	idx := 0
	tmp := make(map[uint64]bool)
	tmp[uint64(1)] = true
	for s.Scan() {
		if b := s.Bytes(); b != nil {
			//逆向切割
			hashKey := hashBytes(b)
			BRANDKEYS[hashKey] = idx
			idx++
		}
	}
	return nil
}

//获取数组的最小项
func getArrayMinOnlineValue(filterkey uint64) (BrandItem, int) {
	minItem := toplist[0]
	minidx := 0
	ilen := len(toplist)
	for i := 1; i < ilen; i++ {
		if toplist[i].OnlineCount < minItem.OnlineCount {
			minItem = toplist[i]
			minidx = i
		}
	}
	return minItem, minidx
}

func sendToTopOnlineList(name, onlineDate []byte, hashKey, nameCount, brandValue uint64) {
	idx, ok := topMap[hashKey]
	if !ok {
		minItem, mIndex := getArrayMinOnlineValue(hashKey)
		isRelpace := false
		if nameCount > minItem.OnlineCount {
			isRelpace = true
		} else if nameCount == minItem.OnlineCount {
			if brandValue > minItem.TotalValue {
				isRelpace = true
			} else if brandValue == minItem.TotalValue {
				if BRANDKEYS[hashKey] < BRANDKEYS[minItem.HashKey] {
					isRelpace = true
				}
			}
		}
		if isRelpace {
			tempKey := minItem.HashKey
			minItem.HashKey = hashKey
			minItem.TotalValue = brandValue
			minItem.OnlineCount = nameCount
			minItem.Name = string(name)
			delete(topMap, tempKey)
			topMap[hashKey] = mIndex
			toplist[mIndex] = minItem
		}
	} else {
		toplist[idx].TotalValue = brandValue
		toplist[idx].OnlineCount = nameCount
	}
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
		if b := s.Bytes(); b != nil {
			//逆向切割
			bs := genSpaceSplit(b)

			{ //限定上架天数最多的品牌
				onlineDate := bs[0] //上线日期
				name := bs[4]

				hashKey := hashBytes(name)
				current_brand_value := BRANDDB[hashKey] + parsebyteToUint64(bs[1])
				BRANDDB[hashKey] = current_brand_value
				//实时处理

				sarray := combineArray(name, onlineDate)
				combineKey := hashBytes(sarray)

				if _, ok := ONLINESMAP[combineKey]; !ok {

					ONLINESMAP[combineKey] = true
					nameCount := ONLINESCOUNT[hashKey] + 1
					ONLINESCOUNT[hashKey] = nameCount
					sendToTopOnlineList(name, onlineDate, hashKey, nameCount, current_brand_value)
				}
			}
		}
	}
	return nil
}

//输出结果
func ListResult() {
	values := make([]BrandItem, TOPNUM)
	for i, item := range toplist {
		values[i] = item
	}
	quickSort(values, 0, len(values)-1)
	for i, item := range values {
		if item.HashKey == 0 {
			continue
		}
		fmt.Printf("(%d) name: %s | value: %d\n", (i + 1), item.Name, item.OnlineCount)
	}
	topMap = nil
	BRANDDB = nil
	fmt.Println("------- finish -------")
}

//快速排序:从大到小
func quickSort(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].OnlineCount
		for i <= j {
			for arr[i].OnlineCount > key {
				i++
			}
			for arr[j].OnlineCount < key {
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
