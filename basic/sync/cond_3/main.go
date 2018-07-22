package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type MyDataBucket struct {
	br     *bytes.Buffer
	gmutex *sync.RWMutex
	rcond  *sync.Cond //读操作需要用到的条件变量
}

func NewDataBucket() *MyDataBucket {
	buf := make([]byte, 0)
	db := &MyDataBucket{
		br:     bytes.NewBuffer(buf),
		gmutex: new(sync.RWMutex),
	}
	db.rcond = sync.NewCond(db.gmutex.RLocker())
	return db
}

func (db *MyDataBucket) Read(i int) {
	db.gmutex.RLock()
	defer db.gmutex.RUnlock()
	var data []byte
	var d byte
	var err error
	for {
		//读取一个字节
		if d, err = db.br.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.rcond.Wait()
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
	}
}

func (db *MyDataBucket) Put(d []byte) (int, error) {
	db.gmutex.Lock()
	defer db.gmutex.Unlock()
	//写入一个数据块
	n, err := db.br.Write(d)
	db.rcond.Broadcast()
	return n, err
}

func main() {
	db := NewDataBucket()

	go db.Read(1)

	go db.Read(2)

	for i := 0; i < 10; i++ {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			db.Put([]byte(d))
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
}
