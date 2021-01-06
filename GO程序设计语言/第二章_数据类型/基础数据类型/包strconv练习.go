package main

import (
	"fmt"
	"strconv"
)

const (
	p1 int = 2 * iota
	p2
	p3
	p4
	p5
)

func main() {
	fmt.Println(strconv.Itoa(123))
	fmt.Println(strconv.Atoi("123"))

	// 十进制转2进制
	fmt.Println(strconv.FormatInt(int64(15), 2))
	fmt.Println(strconv.FormatInt(int64(15), 16))


}
