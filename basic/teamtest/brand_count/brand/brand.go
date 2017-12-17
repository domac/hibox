package brand

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

var brand_db = make(map[uint64][]byte)
var TARGET_STORE = []byte("VIP_NH")

//数据文件读入处理
func ReadAndHandle(brand_db string, dataFile string) error {
	println("------- start -------")
	log.Printf("brand db: %s\n", brand_db)
	log.Printf("dataFile: %s\n", dataFile)

	f, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		if b := s.Bytes(); b != nil {
			bs := genSpaceSplit(b)

			{ //限定上架年份和仓库
				age := parsebyteToUint64(bs[0][2:4])
				if bytes.Compare(bs[2], TARGET_STORE) == 0 && (age >= 11 && age < 17) {
					fmt.Printf("%s\n", bs)
				}
			}
		}
	}
	return nil
}

//输出结果
func ListResult() {
	println("------- finish -------")
}
