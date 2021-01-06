package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buf := bytes.NewBufferString("hello")
	buf.WriteString(" world")
	buf.WriteString(" 腾讯")
	fmt.Println(buf.String())

	fmt.Println("--------------------------")
	for {
		r, _, err := buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n\n读取完毕\n\n")
				break
			}
			panic("读取异常")
		}

		fmt.Printf("%c", r)
	}

	fmt.Println("通过ReadRune()读取，读取完毕后会清空字符串缓冲区:[", buf.String(), "]")
}
