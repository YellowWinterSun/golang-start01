package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}

	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	// 这种初始化方式相当于，初始化一个3容量的string数组，其中0和2位置分别赋值first、third 而其他位置是默认值
	var strs [3]string = [...]string{0: "first", 2: "third"}
	for _, v := range strs {
		fmt.Println(v)
	}
}
