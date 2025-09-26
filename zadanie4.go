package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"	
)

var votes int 
var mutex sync.Mutex

func gen_votes(numbers int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numbers; i++ {
		var count_votes = rand.Intn(10) + 1
		mutex.Lock()
		votes += count_votes
		mutex.Unlock()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10))) 
	}
}

func main() {
	var a sync.WaitGroup
	var number_votes = 10
	a.Add(number_votes)
	for i := 0; i < number_votes; i++ {
		go func() {
			gen_votes(10)
			a.Done()
		}()
	}
	a.Wait()
	fmt.Println("Общее число голосов:", votes)
} 
