package main 

import (
	"fmt"
	"sync"
)

var buffer []string
var mutex sync.Mutex

func logs(message string) {
	mutex.Lock()
	defer mutex.Unlock()
	buffer = append(buffer, message)
	fmt.Println(message)
}

func main() {
	var a sync.WaitGroup
	for i:=0; i < 10; i++ {
		a.Add(1)
		go func (i int) {
			logs(fmt.Sprintf("Запись в журнал %d", i))
			a.Done()
		}(i)
		
	}
	a.Wait()
}