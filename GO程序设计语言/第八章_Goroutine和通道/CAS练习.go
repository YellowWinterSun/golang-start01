package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var num *int32 = new(int32)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i <= 10000; i++ {

			for {
				old := *num
				if ok := atomic.CompareAndSwapInt32(num, old, old+1); ok {
					break
				}
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 10000; i++ {
			for {
				old := *num
				if ok := atomic.CompareAndSwapInt32(num, old, old+1); ok {
					break
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("结果：", *num)
}
