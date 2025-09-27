package main

import (
	"fmt"
	"sync"
)

type metrics struct {
	suc_counter int
	er_counter  int
	mutex       sync.Mutex
}

func new_collector() *metrics {
	return &metrics{}
}

func (metric *metrics) suc_increment() {
	metric.mutex.Lock()
	defer metric.mutex.Unlock()
	metric.suc_counter++
}

func (metric *metrics) error_increment() {
	metric.mutex.Lock()
	defer metric.mutex.Unlock()
	metric.er_counter++
}

func (metric *metrics) report() {
	metric.mutex.Lock()
	defer metric.mutex.Unlock()
	fmt.Printf("Успешных запросов: %d\nОшибок: %d\n", metric.suc_counter, metric.er_counter)
}

func main() {
	var collector = new_collector()
	collector.suc_increment()
	collector.error_increment()
	collector.report()
}

//work