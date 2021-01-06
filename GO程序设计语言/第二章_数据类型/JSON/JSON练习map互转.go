package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := make(map[string]interface{})
	map1["1"] = "hello"
	map1["2"] = "world"
	map1["3"] = 0

	// map -> json
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		panic("map1转成json失败")
	}
	fmt.Println(string(jsonStr))

	// json -> map
	map2 := make(map[string]interface{})
	err = json.Unmarshal(jsonStr, &map2)
	if err != nil {
		panic("json转map异常")
	}

	fmt.Println("map2:\t", map2)

}
