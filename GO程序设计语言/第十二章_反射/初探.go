package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// TypeOf 返回动态类型  ValueOf 返回动态值部分
	var t int = 10
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.ValueOf(t))

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // TypeOf返回具体动态类型，因此是*os.File，而不是io.Writer
	fmt.Println(reflect.ValueOf(w))

	// reflect.Value用法
	tValue := reflect.ValueOf(t)     // 反射变量t的动态值
	tInterface := tValue.Interface() // 动态值可以返回interface{}
	tAfter := tInterface.(int)       // interface{}变量可以用类型转化，转回来
	fmt.Println(tAfter)

	//
	fmt.Println("----------------")
	var z interface{} = "hello"
	zValue := reflect.ValueOf(z)
	fmt.Println(reflect.TypeOf(z)) // 反应的是变量z的 动态类型部分：string
	fmt.Println(zValue.Type())     // 他的动态值的具体类型: string

}
