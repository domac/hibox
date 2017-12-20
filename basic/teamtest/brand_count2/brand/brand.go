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
	ONLINEDB  = make(map[uint64]int, 0)
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
		if b := s.Bytes(); b != nil {
			//逆向切割
			hashKey := hashBytes(b)
			BRANDKEYS[hashKey] = idx
			idx++
		}
	}
	keysLen := idx
	BRANDDB = make([]int, keysLen, keysLen)
	ONLINEDB = make(map[uint64]int, keysLen*3)

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
		index1 := lasIndexN(b, 9, 32)
		index2 := lasIndexIdx(b, index1, 32)
		index3 := lasIndexIdx(b, index2, 32)
		index4 := lasIndexIdx(b, index3, 32)

		//基础数据
		name := b[:index4]
		hashKey := hashBytes(name)
		if xh, ok := BRANDKEYS[hashKey]; ok {
			onlineDate := b[index1+1:]
			price := b[index2+1 : index1]
			combineHashHey := combinehashBytes(onlineDate, xh)

			currentValue := BRANDDB[xh] + parsebyteToInt(price)
			BRANDDB[xh] = currentValue

			ONLINEDB[combineHashHey] = 1

			updateTopList(name, hashKey, combineHashHey, xh, currentValue)
		}
	}

	return nil
}

func updateTopList(name []byte, hashKey, combineHashHey uint64, xh, currentValue int) {
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

		}

	} else {
		compareTopList()
	}
}

func compareTopList() {

}

//输出结果
func ListResult() {
	fmt.Println("------- finish -------")
}
