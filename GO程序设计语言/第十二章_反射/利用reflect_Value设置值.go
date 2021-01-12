package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	xValue := reflect.ValueOf(&x)
	xValueElem := xValue.Elem()
	px := xValueElem.Addr().Interface().(*int)
	*px = 3
	fmt.Println(*px)
	fmt.Println(xValue.CanAddr())
	fmt.Println(xValueElem.CanAddr())
	fmt.Println(xValue.CanSet())
	fmt.Println(xValueElem.CanSet())

	xValueElem.Set(reflect.ValueOf(5))
	fmt.Println(x)

	xValueElem.SetInt(12)
	fmt.Println(x)
}
