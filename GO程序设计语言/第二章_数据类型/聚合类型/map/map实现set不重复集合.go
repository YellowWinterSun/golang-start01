package main

import "fmt"

// 在GO中没有Set数据结构提供，所以一般用map代替
func main() {
	set := make(map[string]bool)

	set["monday"] = true
	fmt.Println(set["monday"])
	fmt.Println(set["friday"])
}
