package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// way1()
	// way2()
	typeSwitchDemo()
}

/*
	1、类型断言，会导致程序崩溃
	语法：x.(T)
*/
func way1() {
	var w io.Writer // 定义接口的变量
	w = os.Stdout   // 赋值 *os.File （它是一个io.Writer）

	// 断言，断言w是不是*os.File
	f := w.(*os.File)
	fmt.Printf("%T\n", f)

	// 断言，w是不是*byte.Buffer
	c := w.(*bytes.Buffer) // 宕机。因为byte.Buffer没有实现io.Writer的方法
	fmt.Printf("%T\n", c)
}

/*
	2、类型断言，不会导致程序崩溃
	语法：x.(T) 但是出参接2个
*/
func way2() {
	var w io.Writer // 定义接口的变量
	w = os.Stdout   // 赋值 *os.File （它是一个io.Writer）

	// 断言，断言w是不是*os.File
	f, ok := w.(*os.File)
	fmt.Printf("%T %v\n", f, ok)

	// 断言，w是不是*byte.Buffer
	c, ok := w.(*bytes.Buffer) // 不会宕机，ok == false
	fmt.Printf("%T %v\n", c, ok)
}

func typeSwitchDemo() {
	fmt.Println(whichType(int(1)))          // 1
	fmt.Println(whichType(int32(1)))        // 1
	fmt.Println(whichType(uint(1)))         // 1
	fmt.Println(whichType(uint8(1)))        // unexpected type uint8 1
	fmt.Println(whichType(true))            // true
	fmt.Println(whichType("hello"))         // hello
	fmt.Println(whichType([]byte("hello"))) // unexpected type []uint8 [104 101 108 108 111]
	fmt.Println(whichType(rune('G')))       // G

}

func whichType(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case rune /* 同时兼容 int32 但是如果下面不用%v | 用%c就只能用于rune，如果是int32的数就会出问题 */ :
		return fmt.Sprintf("%v", x)
	case int, uint:
		return fmt.Sprintf("%d", x)
	case string:
		return fmt.Sprintf("%s", x)
	case bool:
		return fmt.Sprintf("%v", x)
	default:
		return fmt.Sprintf("unexpected type %T %v", x, x)
	}
}
