package main

import "fmt"

func double() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}

func myFun1() int {
	return 1
}

func main() {
	// todo 尝试1
	var f1 func() int
	f1 = double()

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

	// todo 函数变量
	f2 := myFun1
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

	// todo
	var ghostFunc func(int) int
	ghostFunc = func(x int) int {
		if x == 1 {
			return 1
		}

		// 如果要在这里引用ghostFunc，就必须要先声明nil的函数变量，再赋值
		return x + ghostFunc(x - 1)
	}

	fmt.Println(ghostFunc(5))	// 5 + 4 + 3 + 2 + 1
}
