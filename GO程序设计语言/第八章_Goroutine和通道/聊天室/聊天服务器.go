package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// 定义聊天室服务端
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 负责广播消息
	go broadcaster()

	for {
		// 当有客户端建立连接时，这里会返回
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// 异步出去处理单独的客户，以便并发的继续接收新的客户
		go handleConn(conn)
	}
}

type client chan<- string // 对外发送消息的通道
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有接收的客户消息
)

func broadcaster() {
	// 维护一个在线的客户map
	clients := make(map[client]bool) // 所有连接的客户端
	for {
		select {
		case msg := <-messages:
			// 把消息广播给所有client
			fmt.Printf("[广播消息]%s\n", msg)
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			fmt.Printf("[客户端加入]\n")
			clients[cli] = true
		case cli := <-leaving:
			fmt.Printf("[客户端退出]\n")
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
