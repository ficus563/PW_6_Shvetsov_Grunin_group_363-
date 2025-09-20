package main

import (
	"fmt"
	"sync"
	"time"
	_"sync/atomic"
	
)

type tasks struct {
	id	int
	name string
}

func main() {
	var wg sync.WaitGroup
	var queue = make(chan tasks, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i<= 5; i++ {
			var task = tasks{id: i, name: fmt.Sprint("Задча №", i)}
			queue <- task
			fmt.Println("Добавлена задача:", task.name)
			time.Sleep(300 * time.Millisecond)
		}
		close(queue)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range queue {
			fmt.Println("Было обработано", task.name)
			time.Sleep(400 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("Все задачи были обработны...")
} 

//Work