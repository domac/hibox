package brand

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

var brand_db = make(map[uint64][]byte)
var TARGET_STORE = []byte("VIP_NH")

//类目参考格式：
//WyzKsCJkn CrOlZWxM
//rfpGUpcTNbYRMDN
//yQGVedxIMRSOVtoJQJdNZ

//数据参考格式 (104575830)
//VcrmXKjrRfNT RISBDLzIjw aNCIkIHlhlgYZjwJmw ROP_HZ 128234034 2017-1-11
//szCOOcbkHyHIqwSrQknnl BguGGGFuYWcQwwykS ROP_CD 273262040 2008-6-1
//mLCdVg JrVirWHJkSGT EewUcdKFRxIdaxG ROP_HZ 504426487 2016-10-7
//IOotaK qPJXV EWAMpuuS AUVEQWtXGneBFjWAdJRlcA ROP_NH 266501671 2011-3-13
//tYXjAeIYBBYC TXSajPTW CtdfXOrcCjSXSFxijJ ROP_HZ 89424310 2014-5-28

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

func handleSubject1(bs [][]byte) {

}

//输出结果
func ListResult() {
	println("------- finish -------")
}
