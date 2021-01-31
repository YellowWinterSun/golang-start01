package main

import (
	"fmt"
	"sync"
	"time"
)

/*
读取ch超时
写入ch超时
*/
func main() {
	ch := make(chan interface{})
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		select {
		case ch <- 1:
			fmt.Println("写入ch成功")
		case <-time.After(500 * time.Millisecond):
			fmt.Println("写入ch超时")
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case x := <-ch:
			fmt.Println("读取ch成功：", x)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("读取ch超时")
		}
	}()

	wg.Wait()
}
