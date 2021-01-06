package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello 世界"
	eachString(str)
	eachString2(str)
	eachString3(str)
}

func eachString(str string) {
	fmt.Println("-------- eachString ----------")
	// len 返回的是字节总数
	fmt.Println("len", len(str))
	// RuneCountInString 返回的是字符总数。 一个字符（如utf-8）可能对应N个字节组成
	fmt.Println("RuneCountInString", utf8.RuneCountInString(str))

	// 一种遍历字符的写法
	fmt.Printf("\n%%d\t%%c\t%%q\t\t%%d\n")
	for i := 0; i < len(str); {
		// r 字符本身   size 占用字节数
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%d\t%c\t%q\t\t%d\n", i, r, r, r)
		i += size
	}
	fmt.Println("--------------------------------")
}

func eachString2(str string) {
	fmt.Println("-------- eachString2 ----------")
	// len 返回的是字节总数
	fmt.Println("len", len(str))
	// RuneCountInString 返回的是字符总数。 一个字符（如utf-8）可能对应N个字节组成
	fmt.Println("RuneCountInString", utf8.RuneCountInString(str))

	// 一种遍历字符的写法
	fmt.Printf("\n%%d\t%%c\t%%q\t\t%%d\n")
	for i, r := range str {
		// r 字符本身   size 占用字节数
		fmt.Printf("%d\t%c\t%q\t\t%d\n", i, r, r, r)
	}
	fmt.Println("--------------------------------")
}

func eachString3(str string) {
	fmt.Println("-------- eachString3 ----------")
	// len 返回的是字节总数
	fmt.Println("len", len(str))
	// RuneCountInString 返回的是字符总数。 一个字符（如utf-8）可能对应N个字节组成
	fmt.Println("RuneCountInString", utf8.RuneCountInString(str))

	// 一种遍历字符的写法
	for _, r := range str {
		// r 字符本身   size 占用字节数
		fmt.Printf("%c", r)
	}
	fmt.Println("")
	fmt.Println("--------------------------------")
}
