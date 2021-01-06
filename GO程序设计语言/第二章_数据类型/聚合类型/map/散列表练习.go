package main

import "fmt"

func main() {
	hashmap := make(map[string]int)

	hashmap["a"] = 1
	hashmap["b"] = 2


	_, ok := hashmap["a"]
	fmt.Println("a 是否存在 ", ok)
	_, ok = hashmap["c"]
	fmt.Println("c 是否存在 ", ok)

	for k, v := range hashmap {
		fmt.Println(k, v)
	}
}
