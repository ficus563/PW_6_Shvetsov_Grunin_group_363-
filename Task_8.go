package main

import (
	"fmt"
	"sync"
)

type pool_worker struct {
	jobs chan string
	a    sync.WaitGroup
}

func new_worker(size int) *pool_worker {
	var pool = &pool_worker{
		jobs: make(chan string, size),
	}
	for i := 0; i < size; i++ {
		go pool.worker()
	}
	return pool
}

func (pool *pool_worker) worker() {
	for job := range pool.jobs {
		fmt.Println("Обработка задачи:", job)
		pool.a.Done()
	}
}

func (pool *pool_worker) mit(job string) {
	pool.a.Add(1)
	pool.jobs <- job
}

func (pool *pool_worker) Wait() {
	close(pool.jobs)
	pool.a.Wait()
}

func main() {
	var pool = new_worker(3)
	for i := 0; i < 10; i++ {
		job := fmt.Sprintf("Задача %d", i+1)
		pool.mit(job)
	}
	pool.Wait()
}

//work