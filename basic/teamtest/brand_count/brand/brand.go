package brand

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

//全局常量
const TOPNUM = 40

//全局变量
var (
	BRANDKEYS       = make(map[uint64]int, 1024*1024*20)
	BRANDDB         = make(map[uint64]uint64, 1024*1024)
	TARGET_STORE    = []byte("VIP_NH")
	CurrentMinValue uint64 //当前最小值
	toplist         [TOPNUM]BrandItem
	topMap          = make(map[uint64]int)
)

//品牌项
type BrandItem struct {
	Name       string
	TotalValue uint64
	HashKey    uint64

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
func getArrayMinValue(filterkey uint64) (BrandItem, int) {
	minItem := toplist[0]
	minidx := 0
	ilen := len(toplist)
	for i := 1; i < ilen; i++ {
		if toplist[i].TotalValue < minItem.TotalValue {
			minItem = toplist[i]
			minidx = i
		}
	}
	return minItem, minidx
}

func sendToTopList(name []byte, hashKey, brandValue uint64) {
	idx, ok := topMap[hashKey]
	if !ok {
		minItem, mIndex := getArrayMinValue(hashKey)

		isRelpace := false

		if brandValue > minItem.TotalValue {
			isRelpace = true
		} else if brandValue == minItem.TotalValue {
			if BRANDKEYS[hashKey] < BRANDKEYS[minItem.HashKey] {
				isRelpace = true
			}
		}
		//取代之前最小值
		if isRelpace {
			tempKey := minItem.HashKey
			minItem.HashKey = hashKey
			minItem.TotalValue = brandValue
			minItem.Name = string(name)
			delete(topMap, tempKey)
			topMap[hashKey] = mIndex
			toplist[mIndex] = minItem
		}
	} else {
		toplist[idx].TotalValue = brandValue
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

			{ //限定上架年份和仓库
				age := parsebyteToUint64(bs[0][2:4])
				name := bs[4]
				if bytes.Compare(bs[2], TARGET_STORE) == 0 && (age >= 11 && age < 17) {
					hashKey := hashBytes(name)
					current_brand_value := BRANDDB[hashKey] + parsebyteToUint64(bs[1])
					BRANDDB[hashKey] = current_brand_value
					sendToTopList(name, hashKey, current_brand_value)
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
		fmt.Printf("(%d) name: %s | value: %d\n", (i + 1), item.Name, item.TotalValue)
	}
	topMap = nil
	BRANDDB = nil
	fmt.Println("------- finish -------")
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
