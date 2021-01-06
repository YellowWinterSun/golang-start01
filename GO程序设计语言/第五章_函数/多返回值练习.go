package main

import "fmt"

func function1() (result string, err string) {
	result = "function1"
	err = "error"
	return
}

func main() {
	result, _ := function1()
	fmt.Println(result)
	fmt.Println(function1())
}
