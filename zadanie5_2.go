package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	value      interface{}
	time_stamp time.Time
}

type TTL struct {
	items map[string]*CacheItem
	mutex sync.Mutex
}

func new_cache() *TTL {
	return &TTL{
		items: make(map[string]*CacheItem),
	}
}

func (cache *TTL) get(key string) (interface{}, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	item, exists := cache.items[key]
	if !exists || item.time_stamp.Before(time.Now()) {
		delete(cache.items, key)
		return nil, false
	}
	item.time_stamp = time.Now().Add(item.time_stamp.Sub(time.Now()))
	return item.value, true
}

func (cache *TTL) set(key string, v interface{}, tl time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.items[key] = &CacheItem{
		value:      v,
		time_stamp: time.Now().Add(tl),
	}
}

func main() {
	var cache = new_cache()
	cache.set("key", "cache_cache", time.Second*5)
	value, ok := cache.get("key")
	fmt.Println(value, ok)
	time.Sleep(time.Second * 6)
	value, ok = cache.get("key")
	fmt.Println(value, ok)
}
