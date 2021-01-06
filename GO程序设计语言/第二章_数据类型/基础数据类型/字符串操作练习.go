package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(basename("a/b/c.go"))	// "c"
	//fmt.Println(basename("c.d.go"))	// "c.d"
	//fmt.Println(basename("a/b/c/filename.txt"))	// "filename"
	//
	//fmt.Println(basename2("a/b/c.go"))	// "c"
	//fmt.Println(basename2("c.d.go"))	// "c.d"
	//fmt.Println(basename2("a/b/c/filename.txt"))	// "filename"

	fmt.Println(comma("1234567"))
}

func basename(s string) string {
	var lastDotIndex int = -1
	// 将最后一个'/'和之前的字符全部舍去
	for i := len(s) - 1; i >= 0; i-- {
		if lastDotIndex < 0 && s[i] == '.' {
			lastDotIndex = i
			continue
		}
		if s[i] == '/' {
			lastDotIndex -= i + 1
			s = s[i + 1:]
			break
		}
	}

	// 保留最后一个'.'之前的内容。意味着去掉文件编码格式，只输出文件名
	if lastDotIndex == -1 {
		return s
	} else {
		return s[:lastDotIndex]
	}
}

/**
	利用 strings 包
 */
func basename2(s string) string {
	// 如果找不到字符，返回 -1 。跟Java一样
	slash := strings.LastIndex(s, "/")
	s = s[slash + 1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[: dot]
	}

	return s
}

/**
	输入: 1234567
	输出：1,234,567
 */
func comma (s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	if n == 0 {
		n = 3
	}

	buf.WriteString(s[:n])
	for n < len(s) {
		buf.WriteByte(',')
		buf.WriteString(s[n: n+3])
		n += 3
	}

	return buf.String()
}