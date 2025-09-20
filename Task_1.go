package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	
)

var count int64

func main() {
	var wg sync.WaitGroup
	var numCount = 13
	
	for i:=0; i < numCount; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Кол-во просмотров: ", atomic.LoadInt64(&count))
} 

//work