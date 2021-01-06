package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 返回 [32]uint8
	c1 := sha256.Sum256([]byte("password"))
	c2 := sha256.Sum256([]byte("password"))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c1)
	fmt.Println(c1 == c2)


	m1 := md5.Sum([]byte("password"))
	m2 := md5.Sum([]byte("Password"))

	fmt.Printf("%x\n", m1)
	fmt.Printf("%x\n", m2)
	fmt.Println(m1 == m2)
	fmt.Printf("%T", m1)
}
