package main

import "fmt"

type Bird struct {
	Name string
}

func (bird *Bird) SetName(name string) {
	bird.Name = name
}

type MyIntArray []int
func (array *MyIntArray) Total() int {
	if array == nil {
		return -1
	}

	var total int = 0
	for _, v := range *array {
		total += v
	}

	return total
}

func main() {
	var bird *Bird = &Bird{"hdy"}
	fmt.Println(*bird)

	bird.SetName("azure")
	fmt.Println((*bird).Name)

	var array MyIntArray = []int{1, 2, 3}
	var nilArray *MyIntArray

	fmt.Println(array.Total(), nilArray.Total())
}
