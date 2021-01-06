package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	//c := [3]int{1, 2}
	d := [2]int{1, 3}
	e := [...]int{1, 2}

	fmt.Println(a == b, a == d)
	fmt.Println(b == e)
	//fmt.Println(a == c)	// 编译错误 a和c属于不同类型

}
