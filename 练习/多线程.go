package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// 2个线程处理异步统计
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// 这里会阻塞住，知道 chan信道有数据传输过来
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
