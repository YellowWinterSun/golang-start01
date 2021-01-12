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
	stu := student{1, "hdy", "ok"}
	stuType := reflect.TypeOf(stu)
	f, ok := stuType.FieldByName("id")
	fmt.Println(ok)
	fmt.Println(f.Name, f.Type)
	fmt.Println(f.Type.Kind() == reflect.Int32)
	fmt.Println(f.Type.Kind() == reflect.Int)
	fmt.Println(stuType.NumField())
	fmt.Println(stuType.NumMethod())

}
