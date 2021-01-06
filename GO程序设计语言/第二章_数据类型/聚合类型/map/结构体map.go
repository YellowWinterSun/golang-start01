package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func (stu Student) String() string {
	return stu.Name
}

func main() {
	hashmap := make(map[Student]Student)

	hashmap[Student{"hdy"}] = Student{"kxx"}
	fmt.Println(hashmap)

	if value, ok := hashmap[Student{"hdy"}]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("找不到hdy")
	}


	for k, v := range hashmap {
		fmt.Printf("%s\t%s\n", k, v)
	}
}
