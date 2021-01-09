package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 接口值有2个概念：动态类型、值
func main() {
	// 定义接口变量
	var w io.Writer
	fmt.Printf("%T\n", w) // nil	动态类型:nil  值:nil

	// os.Stdout返回的*File 是 io.Writer的一个实现
	w = os.Stdout
	fmt.Printf("%T\n", w) // *os.File	  动态类型：*os.File	值：指向具体存放数据的地址

	// new()函数返回 类型的指针，会为类型进行创建开辟内存空间，然后返回指针
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // *bytes.Buffer 动态类型：*bytes.Buffer	值：指向具体内存

	w = nil
	fmt.Printf("%T\n", w) // nil nil

	/**
	2、空指针的非空接口
	*/
	//errorDemo(false)
	goodDemo(false)
}

func errorDemo(debug bool) {
	var buf *bytes.Buffer // nil
	if debug {
		buf = new(bytes.Buffer)
	}

	// 当debug是false 会导致宕机	，这里虽然buf是nil，但是传进去后。动态类型是*bytes.Buffer，动态值是nil。
	// 这样导致的结果是，里面的out != nil 会成立，从而导致宕机
	f(buf)
}

// errorDemo对应的正确写法
func goodDemo(debug bool) {
	// todo 跟errorDemo的本质区别在这里
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}

	// 这里就不会报错
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	} else {
		fmt.Println("调用f时，out == nil")
	}
}
