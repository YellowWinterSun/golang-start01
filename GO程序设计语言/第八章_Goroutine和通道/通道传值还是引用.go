package main

import (
	"fmt"
	"time"
)

type chanDTO struct {
	value string
}

// 这个例子可以看出，channal中传递数据时，普通类型和普通结构体，都是值传递
func main() {
	ch := make(chan chanDTO)

	go func() {
		v := chanDTO{"hello"}
		fmt.Printf("t1 %p\n", &v)
		ch <- v
		time.Sleep(2 * time.Second)
		fmt.Println("t1", v)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		v := <-ch
		fmt.Printf("t2 %p\n", &v)

		v.value = "world"
		fmt.Println("t2", v)
	}()

	for {

	}
}
