package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	/*
		todo 1 JSON转struct
	 */
	jsonStr := `{"name":"hdy", "age": 25}`
	var hdy People
	if err := json.Unmarshal([]byte(jsonStr), &hdy); err != nil {
		panic("jsonStr转struct失败")
	}
	fmt.Println(hdy)

	fmt.Println("------- 2 ---------")
	/*
		todo 2 struct转JSON
	 */
	azure := People{"azure", 26}
	jsonByteStr, err := json.MarshalIndent(azure, "", "\t")
	if err != nil {
		panic("struct转json失败")
	}

	fmt.Println("azure转JSON后：\n", string(jsonByteStr))

	// 再转回来
	azure2 := People{}
	if err = json.Unmarshal(jsonByteStr, &azure2); err != nil {
		panic("异常")
	}

	fmt.Println("azure2:\n", azure2)
}
