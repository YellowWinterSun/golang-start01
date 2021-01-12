package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id     int    `http:"ok" h:"this is h" b:"this is b"`
	name   string `NAME`
	remark string `REMARK`
}

func main() {
	stu1 := student{1, "hey", "good"}
	var i interface{}
	i = &stu1

	elem := reflect.ValueOf(i).Elem()
	sf := elem.Type().Field(0) // StructField
	fmt.Println(sf.Tag)
	fmt.Println(sf.Tag.Get("http"))
	fmt.Println(sf.Tag.Get("h"))
	fmt.Println(sf.Tag.Get("b"))
	v := elem.Field(0) // Value
	fmt.Println(sf, v)
}
