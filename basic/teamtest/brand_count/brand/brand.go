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
	//BRANDDB      = make(map[uint64]int, 1024*1024)
	BRANDDB = []int{}
	toplist [TOPNUM]uint64
	topMap  = make(map[uint64]int)
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
	println(len(BRANDDB))
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
		if age != 54 && age != 53 && age != 52 && age != 51 && age != 50 {
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
			BRANDDB[xh] += parsebyteToInt(price)
			updateTopList(name, hashKey, xh)
		}

		// onlineDate := b[index1+1:]
		// fmt.Printf("%s\n", onlineDate)
		// fmt.Printf("%s\n", price)
		// fmt.Printf("%s\n", name)
		// println("-----------")
	}
	return nil
}

func updateTopList(name []byte, hashKey uint64, xh int) {

}

//输出结果
func ListResult() {
	fmt.Println("------- finish -------")
}
