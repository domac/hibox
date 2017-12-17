package main

import (
	b "./brand"
	"log"
	"os"
	"time"
)

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

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("please input the data file path")
		os.Exit(2)
	}
	start := time.Now()
	err := b.ReadAndHandle(args[len(args)-2], args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}
	elapsed := time.Now().Sub(start)
	log.Printf("total elapsed time: %f seconds", elapsed.Seconds())
	b.ListResult()
}
