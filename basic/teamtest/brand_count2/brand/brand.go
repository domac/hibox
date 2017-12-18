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
	BRANDKEYS = make(map[uint64]int, 1024*1024*20)
	BRANDDB   = make(map[uint64]uint64, 1024*1024)
	topMap    = make(map[uint64]int)
)

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
			fmt.Printf("%s\n", bs)
		}
	}
	return nil
}

//输出结果
func ListResult() {
	fmt.Println("------- finish -------")
}
