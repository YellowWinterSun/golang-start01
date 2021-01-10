package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		fmt.Printf("ERROR: %v", err.Error())
		return
	}

	// 客户端读取聊天室消息行为
	go func() {
		//buffer := make([]byte, 1024)
		//for {
		//	n, err := conn.Read(buffer)
		//	if err != nil {
		//		break
		//	}
		//
		//	fmt.Printf("%s\n", string(buffer[0:n]))
		//}

		reader := bufio.NewReader(conn)
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			fmt.Printf("%s", str)
		}

	}()

	// 客户端发送行为
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 回车表示一条消息
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取控制台输入异常, %s\n", err)
			break
		}

		if strings.TrimSpace(input) == "quit" {
			break
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Printf("写入异常%v\n", err)
			break
		}
	}
}
