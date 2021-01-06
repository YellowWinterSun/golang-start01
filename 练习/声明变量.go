package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 普通声明变量  var 变量名 数据类型
	var name string = "Hdy"
	var number int32 = 1000
	var money float64 = 1000.01
	var isMan bool = true
	var tags []string

	const CONST_NAME = "FUCK"

	for i:=0;i<5;i++ {
		tags=append(tags, strconv.Itoa(i))
	}

	fmt.Println(name, number, money, isMan, CONST_NAME)

	fmt.Println(tags[1])
}
