package main 

import (
	"fmt"
	"sync"
	
)

type product struct {
	name string
	stock int
	mutex sync.Mutex 
}

func (p * product) sell(amount int) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if amount > p.stock {
		return fmt.Errorf("Недостаточно товара на складе %s", p.name)
		
		
	} 
	p.stock -= amount
	return nil


}

func(p * product) restock(amount int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.stock += amount
}

func main() {
	var products = &product{name: "Штаны", stock: 300}
	var err = products.sell(50)
	if err != nil {
		fmt.Println(err)
	}
	products.restock(60)
	fmt.Println(products.stock)
}