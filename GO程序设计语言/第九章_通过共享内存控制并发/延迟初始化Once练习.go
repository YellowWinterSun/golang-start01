package main

import (
	"fmt"
	"sync"
)

var initOnce sync.Once
var myMap map[string]string

func initMyMap() {
	fmt.Println("init...")
	myMap = make(map[string]string)
	myMap["A"] = "a"
	myMap["B"] = "b"
	myMap["C"] = "c"
	myMap["D"] = "d"
}

func getMyMap(key string) (v string, ok bool) {
	initOnce.Do(initMyMap)
	v, ok = myMap[key]
	return
}

func main() {
	fmt.Println(getMyMap("A"))
	fmt.Println(getMyMap("Z"))
}
