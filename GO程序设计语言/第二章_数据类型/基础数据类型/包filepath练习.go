package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	str := "a/b/c.go"

	// 获取文件名
	fmt.Println(filepath.Base(str))

}
