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

	dataList  = []int{}
	namedList = []string{}
)

type BrandItem struct {
	Name    string
	HashKey uint64
	xh      int

	DateCount  int
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
	dataList = make([]int, keysLen, keysLen)
	namedList = make([]string, keysLen, keysLen)

	for i := 0; i < keysLen; i++ {
		BRANDDB[i] = 0
		dataList[i] = 0
		namedList[i] = ""
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
		s.Bytes()

	}
	return nil
}

//输出结果
func ListResult() {

	fmt.Println("------- finish -------")
}
