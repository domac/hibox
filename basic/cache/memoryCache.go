package cache

import (
	"errors"
	"log"
	"sync"
	"time"
)

type MemoryCache struct {
	sync.RWMutex
	selfCheckInterval time.Duration
	items             map[string]*MemItem
	Every             int
}

type MemItem struct {
	val           interface{}
	createdTime   time.Time
	expireTimeout time.Duration
}

//缓存项是否过期
func (item *MemItem) isExpire() bool {
	if item.expireTimeout < 0 {
		return false
	}
	return time.Now().Sub(item.createdTime) > item.expireTimeout
}

func NewMemoryCache(interval int) (*MemoryCache, error) {
	memoryCahce := &MemoryCache{
		items: make(map[string]*MemItem),
	}
	err := memoryCahce.StartAndGC(interval)
	if err != nil {
		return nil, err
	}
	return memoryCahce, nil
}

func (mc *MemoryCache) GetMulti(keys []string) []interface{} {
	var rc []interface{}
	for _, key := range keys {
		rc = append(rc, mc.get0(key))
	}
	return rc
}

func (mc *MemoryCache) Get(key string) interface{} {
	mc.RLock()
	defer mc.RUnlock()
	return mc.get0(key)
}

func (mc *MemoryCache) get0(key string) interface{} {
	if itm, ok := mc.items[key]; ok {
		if itm.isExpire() {
			return nil
		}
		return itm.val
	}
	return nil
}

func (mc *MemoryCache) Put(key string, value interface{}, expireTimeout time.Duration) error {
	mc.Lock()
	mc.Unlock()
	mc.items[key] = &MemItem{
		val:           value,
		createdTime:   time.Now(),
		expireTimeout: expireTimeout,
	}
	return nil
}

func (mc *MemoryCache) Delete(name string) error {
	mc.Lock()
	defer mc.Unlock()
	if _, ok := mc.items[name]; !ok {
		return errors.New("key not exist")
	}
	delete(mc.items, name)
	if _, ok := mc.items[name]; ok {
		return errors.New("delete key error")
	}
	return nil
}

func (mc *MemoryCache) IsExist(name string) bool {
	mc.RLock()
	defer mc.RUnlock()
	if v, ok := mc.items[name]; ok {
		return !v.isExpire()
	}
	return false
}

func (mc *MemoryCache) ClearAll() error {
	mc.Lock()
	defer mc.Unlock()
	mc.items = make(map[string]*MemItem)
	return nil
}

func (mc *MemoryCache) StartAndGC(interval int) error {
	if interval < 0 {
		return errors.New("start and gc error: interval must more than or equal zero")
	}
	mc.selfCheckInterval = time.Duration(interval) * time.Second
	mc.Every = interval
	go mc.vaccuum()
	return nil
}

//自检函数
func (mc *MemoryCache) vaccuum() {
	if mc.Every < 1 {
		return
	}
	timer := time.NewTimer(mc.selfCheckInterval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			if mc.items == nil {
				log.Println("memory items is null, exit vaccuum")
				return
			}
			log.Println("cache self check ...")
			if keys := mc.expiredKeys(); len(keys) != 0 {
				mc.clearItems(keys)
			}
			timer.Reset(mc.selfCheckInterval)
		}
	}
}

func (mc *MemoryCache) expiredKeys() (keys []string) {
	mc.RLock()
	defer mc.RUnlock()
	for key, item := range mc.items {
		if item.isExpire() {
			keys = append(keys, key)
		}
	}
	return
}

func (mc *MemoryCache) clearItems(keys []string) {
	mc.Lock()
	defer mc.Unlock()
	for _, key := range keys {
		delete(mc.items, key)
	}
}
