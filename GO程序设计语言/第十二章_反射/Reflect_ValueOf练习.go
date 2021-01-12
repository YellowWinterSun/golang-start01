package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 1
	ptr := &x
	fmt.Printf("地址：%p\n", &x)  // 0xc00000a0b0
	fmt.Printf("地址：%p\n", ptr) // 0xc00000a0b0
	fmt.Println(Any(x))
	fmt.Println(Any(&x))   // 0xc00000a0b0
	fmt.Println(Any(ptr))  // 0xc00000a0b0
	fmt.Println(Any(&ptr)) // 0xc000006028 指针的指针
	fmt.Println(Any([3]int{1, 2, 3}))
	fmt.Println(Any([]int{1, 2, 3}))
	fmt.Println(Any(make(map[string]string)))
	fmt.Println(Any("hello"))

	fmt.Println("--------- 指针 -------------")
	ptrValue := reflect.ValueOf(ptr)
	fmt.Println(ptrValue.Kind() == reflect.Ptr)
	fmt.Println(ptrValue.IsNil(), ptrValue.IsValid(), ptrValue.IsZero())
	ptrValueValue := ptrValue.Elem() // 获取指针指向的变量的Value
	fmt.Println("指针的类型:", ptrValue.Type())
	fmt.Println("指针指向的地址的类型", ptrValueValue.Type())

	fmt.Println("----------- array ------------")

	// 数组 reflect.Value Len() Cap() Index()
	arrayValue := reflect.ValueOf([3]int{4, 5, 6})
	fmt.Println(arrayValue.Kind().String()) // array
	for i := 0; i < arrayValue.Len(); i++ {
		fmt.Println(arrayValue.Index(i).Int())
	}

	sliceValue := reflect.ValueOf(make([]int, 0, 3))
	fmt.Println(sliceValue.Kind().String()) // slice
	fmt.Println(sliceValue.Len())
	fmt.Println(sliceValue.Cap())

	fmt.Println("-------- map -----------")
	// map反射
	myMap := map[string]int{
		"hello": 1,
		"hi":    2,
	}
	mapValue := reflect.ValueOf(myMap)
	fmt.Println(mapValue.Kind() == reflect.Map)
	for _, key := range mapValue.MapKeys() {
		fmt.Printf("Key:[%s] value:[%v]\n", key.String(), mapValue.MapIndex(key))
	}
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// reflect.Value 可以让你在不知道interface{}是什么类型时，进行协助判断
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Interface:
		if v.IsNil() {
			return "nil interface{}"
		} else {
			return "interface{}"
		}
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		// 输出类型和十六进制引用地址
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}
