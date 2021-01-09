package main

import "fmt"

func main() {
	var any interface{}
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = []byte("good")
	fmt.Printf("%s\n", any)

	any = rune('世')
	fmt.Printf("%c\n", any)

}
