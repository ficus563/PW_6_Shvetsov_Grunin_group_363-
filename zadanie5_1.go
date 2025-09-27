package main

import (
	"fmt"
	"sync"
	"time"
)

type CachedItem struct {
	Value     interface{}
	Timestamp time.Time
}

type TTLCache struct {
	items map[string]*CachedItem
	mu    sync.Mutex
}

func NewTTLCache() *TTLCache {
	return &TTLCache{
		items: make(map[string]*CachedItem),
	}
}

func (cache *TTLCache) Get(key string) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	item, exists := cache.items[key]
	if !exists || item.Timestamp.Before(time.Now()) {
		delete(cache.items, key)
		return nil, false
	}
	return item.Value, true
}

func (cache *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.items[key] = &CachedItem{
		Value:     value,
		Timestamp: time.Now().Add(ttl),
	}
}

func main() {
	cache := NewTTLCache()
	cache.Set("key1", "value1", time.Second*5)
	value, ok := cache.Get("key1")
	fmt.Println(value, ok)
	time.Sleep(time.Second * 6)
	value, ok = cache.Get("key1")
	fmt.Println(value, ok)
}
