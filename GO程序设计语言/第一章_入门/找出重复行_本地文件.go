package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := "D:\\mymq.txt"

	countsMap := make(map[string]int)

	// 读取文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
	}

	// 将文件内容，按照换行符进行分割。分割出一个 数组
	fmt.Println("文件内容\n", string(data))
	for _, line := range strings.Split(string(data), " ") {
		countsMap[line]++
	}

	fmt.Println("------文件内容结果--------")
	for k, v := range countsMap {
		fmt.Println(k, v)
	}
}
