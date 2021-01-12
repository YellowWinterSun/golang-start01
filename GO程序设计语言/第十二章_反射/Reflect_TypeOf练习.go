package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id     int32
	name   string
	remark string
}

func (s *student) SetRemark(remark string) {
	s.remark = remark
}

func (s *student) Call() string {
	return s.name
}

func main() {
	stu1 := student{1, "hdy", ""}
	stuType := reflect.TypeOf(stu1)
	fmt.Println(stuType, stuType.Name(), stuType.NumField())

	for i := 0; i < stuType.NumField(); i++ {
		fmt.Printf("%s.%s\t%s\n", stuType.Name(), stuType.Field(i).Name, stuType.Field(i).Type)
	}

	for i := 0; i < stuType.NumMethod(); i++ {
		fmt.Printf("%s.%s\t%s\n", stuType.Name(), stuType.Method(i).Name, stuType.Method(i).Type)
	}
}
