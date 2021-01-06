package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	str := "D:\\mymq.txt"

	f, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	/*
	// todo 按字节读的方式

	fmt.Println("----------------")
	// 一次读4个字节的字节数组
	buf := make([]byte, 4)
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		fmt.Printf("%s", buf[:n])

		if err != nil {
			if err == io.EOF {
				fmt.Println("\n\n\n完成读取\n")
				return
			}
			fmt.Println(err)
			return
		}
	}
	 */


	// todo 2 按字符读
	/*
	bfRd := bufio.NewReader(f)
	for {
		r, _, err := bfRd.ReadRune()
		fmt.Printf("%c", r)

		if err != nil {
			if err == io.EOF {
				fmt.Println("\n\n\n完成读取\n")
				return
			}
			fmt.Println(err)
			return
		}
	}
	 */

	// todo 3 按行读
	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		fmt.Printf("%s", line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n\n\n完成读取\n")
				return
			}
			fmt.Println(err)
			return
		}
	}
}
