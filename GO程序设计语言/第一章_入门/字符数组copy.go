package main

import "fmt"

func main() {
	buf1 := []byte("hello _____!")
	buf2 := []byte("world")

	copy(buf1[6:], buf2)

	fmt.Println(string(buf1))
	fmt.Println(string(buf2))
}
