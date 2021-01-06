package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// strings 操作 字符串
	fmt.Println(strings.Contains("hello world", "hello"))
	fmt.Println(strings.Count("hello", "l"))
	fmt.Println(strings.Fields("hello world"))
	fmt.Println(strings.Index("hello world", "world"))
	fmt.Println(strings.Join([]string{"hello", "world"}, ","))

	// bytes操作 byte数组
	fmt.Println(bytes.Contains([]byte("hello"), []byte("h")))
}
