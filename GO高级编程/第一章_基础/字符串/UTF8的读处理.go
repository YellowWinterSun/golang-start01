package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	/*
		遇到UTF-8的中文字符，会跳着读
		0 世
		3 界
		6 a
		7 b
		8 c
	*/
	for i, c := range "世界abc" {
		fmt.Printf("%d %c\n", i, c)
	}

	/*
		%v 和 %#v 一样能输出，但是加了#会给两端加上双引号
		%v   世界A
		%#v "世界A"
	*/
	fmt.Printf("%#v\n", string([]rune{'世', '界', 'A'}))

	s := bytes2str([]byte("世界hello"))
	fmt.Println(s)
}

// 模拟 string(bytes) 的实现
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}
