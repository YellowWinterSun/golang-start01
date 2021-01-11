package main

import (
	"fmt"
	"sync"
)

var (
	lock    sync.Mutex
	balance int
)

func deposit(amount int) {
	lock.Lock()
	defer lock.Unlock()
	balance += amount
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	goFunc := func() {
		for i := 1; i <= 10000; i++ {
			deposit(1)
		}
		wg.Done()
	}

	go goFunc()
	go goFunc()
	go goFunc()

	wg.Wait()
	fmt.Println(balance)
}
