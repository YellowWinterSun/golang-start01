package main

import (
	"fmt"
	"path"
)

func main() {
	str := "a/b/c.go"
	// 文件全称 c.go
	fmt.Println(path.Base(str))
	fmt.Println(path.Clean("./" + str))
	// 获取全目录路径 a/b
	fmt.Println(path.Dir(str))
	// 获取文件格式 .go
	fmt.Println(path.Ext(str))

	fmt.Println(path.IsAbs(str))

	fmt.Println(path.Split(str))
}
