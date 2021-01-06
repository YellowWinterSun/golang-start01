package main

import "fmt"

func main() {
	// 原始数组
	array := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array, len(array), cap(array))

	// slice
	slice1 := array[0:1]
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = array[2:]
	fmt.Println(slice1, len(slice1), cap(slice1))

	var s []int = make([]int, 0, 10)
	fmt.Println(s == nil)
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
