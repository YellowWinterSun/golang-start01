package main

import (
	"fmt"
	"time"
)

// 在GO中称之为竟态：
func main() {
	demo2()
}

// 并发不安全的例子
func demo1() {
	var balance int
	deposit := func(amount int) {
		balance += amount
	}

	getBalance := func() int {
		return balance
	}

	go func() {
		// Alice 存钱
		deposit(200)
		fmt.Println(getBalance())
	}()

	go func() {
		// Bob 存钱
		deposit(100)
		fmt.Println(getBalance())
	}()

	// 极端并发情况下，两个Alice和Bob的代码交替执行，导致最终银行balance的金额不是300
}

// 线程安全的第一种写法：利用channal编程
func demo2() {
	// 利用通道来实现串行
	var depositChannal = make(chan int) // 存款通道
	var balanceChannal = make(chan int) // 读取余额通道

	deposit := func(amount int) {
		depositChannal <- amount
	}

	getBalance := func() int {
		return <-balanceChannal
	}

	// 专门的单线程处理通道
	go func() {
		var balance int
		for {
			select {
			case amount := <-depositChannal:
				balance += amount
			case balanceChannal <- balance:
			}
		}
	}()

	go func() {
		fmt.Println("t1", getBalance())
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("t1", getBalance())
		}

	}()

	go func() {
		time.Sleep(1 * time.Second)
		deposit(100)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		deposit(200)
	}()

	for {

	}
}

// 并发控制：channal实现互斥的写法
func demo3() {
	var (
		sema    = make(chan struct{}, 1)
		balance int
	)

	deposit := func(amount int) {
		sema <- struct{}{} // 获取令牌，如果已经被别人占用，那么会阻塞等待
		balance += amount
		<-sema // 释放令牌
	}

	getBalance := func() int {
		sema <- struct{}{}
		b := balance
		<-sema
		return b
	}
}
