package main

import (
	"encoding/json"
	"fmt"
)

type Goods struct {
	Name string
	Price float64
	Number uint32
}

type Goods2 struct {
	Name string
	Price float64
	Number uint32
}


type Order struct {
	OrderNo string
	OrderUsername string
	OrderGoods Goods
}

type HighGoods struct {
	HighId int
	A Goods
	B Goods2
}

func main() {
	goods1 := Goods{"apple", 5.2, 10}

	order1 := Order{
		"no1",
		"hdy",
		goods1,
	}

	// 格式化输出json
	jsonData, _ := json.MarshalIndent(order1, "", "\t")
	fmt.Println(string(jsonData))

	/**
		2
	 */
	hg1 := HighGoods{1, goods1, Goods2{"cpu", 1000.0, 10}}
	fmt.Println("json前:\n", hg1)
	// 只能格式化可导的字段
	jsonData, _ = json.Marshal(hg1)
	fmt.Println(string(jsonData))

	// 反序列化
	var hg2 HighGoods
	if err := json.Unmarshal(jsonData, &hg2); err != nil {
		fmt.Printf("JSON转化失败: %s\n", err)
	}

	fmt.Println("json后:\n", hg2)
	fmt.Println(hg1 == hg2)
}
