package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 定义查重用的 map
	countsMap := make(map[string]int)

	// 定义控制台输入
	input := bufio.NewScanner(os.Stdin)

	// 行遍历
	for input.Scan() {
		text := input.Text()
		if text == "finish" {
			break
		}

		countsMap[text]++
	}

	fmt.Printf("%10s\t%10s\n", "key", "value")
	for k, v := range countsMap {
		fmt.Printf("%10s\t%10d\n", k, v)
	}
}
