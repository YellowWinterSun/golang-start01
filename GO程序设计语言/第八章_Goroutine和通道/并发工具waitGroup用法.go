package main

import (
	"fmt"
	"sync"
	"time"
)

const goroutineNum = 5

/**
sync.WaitGroup 类似于 Java juc包的CountDownLatch
sync.WaitGroup通过信号量来控制go协程的等待和工作状态。而Java AQS则是通过控制thread的park unpark从而控制状态
*/
func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go func(index int) {
			time.Sleep(time.Duration(index) * time.Second)
			fmt.Printf("go-%d 结束\n", index)
			waitGroup.Done()
		}(i)
	}

	fmt.Println("主程序等待")
	waitGroup.Wait()
	fmt.Println("--主程序结束等待--")
}
