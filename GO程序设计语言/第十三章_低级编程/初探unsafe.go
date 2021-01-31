package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// SizeOf返回给定参数的在内存中占用的字节长度
	fmt.Println(unsafe.Sizeof(int32(0)))
	fmt.Println(unsafe.Sizeof(float32(0)))

}
