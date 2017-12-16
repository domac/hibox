package brand

import (
	"bufio"
	"errors"
	"log"
	"os"
)

var brand_db = make(map[uint64][]byte)

//类目参考格式：
//WyzKsCJkn CrOlZWxM
//rfpGUpcTNbYRMDN
//yQGVedxIMRSOVtoJQJdNZ

//数据参考格式
//VcrmXKjrRfNT RISBDLzIjw aNCIkIHlhlgYZjwJmw ROP_HZ 128234034 2017-1-11
//szCOOcbkHyHIqwSrQknnl BguGGGFuYWcQwwykS ROP_CD 273262040 2008-6-1
//mLCdVg JrVirWHJkSGT EewUcdKFRxIdaxG ROP_HZ 504426487 2016-10-7
//IOotaK qPJXV EWAMpuuS AUVEQWtXGneBFjWAdJRlcA ROP_NH 266501671 2011-3-13
//tYXjAeIYBBYC TXSajPTW CtdfXOrcCjSXSFxijJ ROP_HZ 89424310 2014-5-28

func ReadAndHandle(brand_db string, dataFile string) error {
	log.Printf("brand db: %s\n", brand_db)
	log.Printf("dataFile: %s\n", dataFile)

	if dataFile == "" {
		return errors.New("data file is null")
	}
	println("------- start -------")
	f, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		if b := s.Bytes(); b != nil {
		}
	}
	return nil
}

//输出结果
func ListResult() {
	println("------- finish -------")
}
