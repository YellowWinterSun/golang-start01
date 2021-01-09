package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	demo3()
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
	tick := time.Tick(1 * time.Second)
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println("倒计时", countdown, "秒")
		select {
		case <-tick:
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
