package main

import (
	"fmt"
	"strings"
)

func main() {
	var str []string = []string{"Hello", "World"}

	fmt.Println(strings.Join(str, " "))
}
