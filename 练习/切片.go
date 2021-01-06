package main

import "fmt"

func main() {
	array := []int{1,2,3,4,5}

	fmt.Println(array)

	arr1 := array[:3]
	fmt.Println(arr1)

	arr2 := arr1[1:]
	fmt.Println(arr2)

	fmt.Println("---------------")
	arr2[0] = -1
	fmt.Println(arr2, array, len(arr2))
}
