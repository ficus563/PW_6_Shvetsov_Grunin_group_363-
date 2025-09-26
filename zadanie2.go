package main

import (
	"fmt"
	"sync"	
)

type cache struct{
	data map[string]string
	mutex sync.RWMutex
}

func new_cache() *cache {
	return &cache{
		data: make(map[string]string),
	}
}

func (c *cache) get(k string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data[k]
}

func (c *cache) set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

func main() {
	var cach = new_cache()
	cach.set("key1", "value1")
	fmt.Println(cach.get("key1"))
} 
