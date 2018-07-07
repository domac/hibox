package cache

import (
	"testing"
	"time"
)

func TestMemoryCahce(t *testing.T) {

	memoryCache, _ := NewMemoryCache(10)

	exprire := time.Second * 5

	memoryCache.Put("domac", "domacli", exprire)
	val := memoryCache.Get("domac").(string)
	println(val)
	if val != "domacli" {
		t.Error("get error")
	}

	time.Sleep(5 * time.Second)
	val0 := memoryCache.Get("domac")
	if val0 != nil {
		t.Error("expire error")
	}

	time.Sleep(300 * time.Second)
}
