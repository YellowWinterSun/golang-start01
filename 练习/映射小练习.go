package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	strArray := strings.Fields(s)

	resultMap := make(map[string]int)
	var temp int
	for _, v := range strArray {
		temp = resultMap[v]
		resultMap[v] = temp + 1
	}

	return resultMap
}

func main() {
	s := "a b c a c"
	fmt.Println(s)
	fmt.Println("---")
	fmt.Println(WordCount(s))
}
