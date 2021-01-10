package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	demo6()
}

// 火箭倒计时DEMO
func demo1() {
	// 定义中止通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字符
		abort <- struct{}{}
	}()

	fmt.Println("倒计时5s后开始..... 输入任意键中止程序")

	select {
	case <-time.After(5 * time.Second):
		// noting to do
	case <-abort:
		fmt.Println("******* 中止 ********")
		return
	}

	fmt.Println("------ 发射成功 --------")
}

// 火箭倒计时DEMO
func demo3() {
	// 定义中止通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字符
		abort <- struct{}{}
	}()

	fmt.Println("倒计时5s后开始..... 输入任意键中止程序")
	//tick := time.Tick(1 * time.Second)	 仅当整个生命周期都需要 tick 时，才使用这种方式，否则无法关闭，导致goroutine泄露
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println("倒计时", countdown, "秒")
		select {
		case <-tick.C:
			// noting to do,just loop
		case <-abort:
			fmt.Println("******* 中止 ********")
			return
		}
	}

	fmt.Println("------ 发射成功 --------")
}

//
func demo2() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	defer close(ch1)
	defer close(ch2)
	defer close(ch3)

	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			s := input.Text()
			switch s {
			case "ch1":
				ch1 <- "ch1 active"
			case "ch2":
				ch2 <- "ch2 active"
			default:
				ch3 <- "exit"
				return
			}
		}
	}()

	for {
		select {
		case x := <-ch1:
			fmt.Printf("来自ch1:[%s]\n", x)
		case x := <-ch2:
			fmt.Printf("来自ch2:[%s]\n", x)
		case <-ch3:
			fmt.Printf("来自ch3，即将中止程序\n")
			goto END
		}
	}

END: // goto mark
	fmt.Println("--------- end ------------")

}

// 生产者，消费者模式
func demo4() {
	ch := make(chan string, 100)

	runable := func(consumerName string, ch chan string) {
		defer fmt.Printf("消费者[%s]结束\n", consumerName)
		for {
			fmt.Printf("<<<<<<<<<<消费者[%s]就绪>>>>>>>>>\n", consumerName)
			select {
			case x, ok := <-ch:
				if !ok {
					return
				}
				fmt.Printf("[%s]拉取到消息:%s\n", consumerName, x)
			}
		}
	}

	go runable("c1", ch)
	go runable("c2", ch)

	input := bufio.NewScanner(os.Stdin)
	for {
		for input.Scan() {
			s := input.Text()
			switch s {
			case "close":
				close(ch)
				goto END
			default:
				ch <- s
			}
		}
	}
END:
	time.Sleep(1 * time.Second)
	fmt.Println("------ 程序结束 ---------")

}

// 非阻塞通信
func demo5() {
	ch := make(chan int)

	//<-ch	//如果直接这样写，会阻塞等待

	// 如果ch没有元素，则马上返回
	select {
	case <-ch:
		fmt.Println("接收到了")
	default:
		fmt.Println("没有东西")
	}

	close(ch)
}

// 控制其他goroutine退出
func demo6() {
	ch := make(chan string, 100)

	runable := func(consumerName string, ch chan string) {
		defer fmt.Printf("消费者[%s]结束\n", consumerName)
		for {
			fmt.Printf("<<<<<<<<<<消费者[%s]就绪>>>>>>>>>\n", consumerName)
			if x, ok := <-ch; ok {
				fmt.Printf("[%s]拉取到消息:%s\n", consumerName, x)
			} else {
				break
			}
		}
	}

	go runable("c1", ch)
	go runable("c2", ch)

	input := bufio.NewScanner(os.Stdin)
	for {
		for input.Scan() {
			s := input.Text()
			switch s {
			case "close":
				close(ch)
				goto END
			default:
				ch <- s
			}
		}
	}
END:
	time.Sleep(1 * time.Second)
	fmt.Println("------ 程序结束 ---------")

}
