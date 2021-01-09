package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for k, v := range db {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	case "/getDollars":
		itemName := req.URL.Query().Get("item")
		if price, ok := db[itemName]; ok {
			fmt.Fprintf(w, "存在商品%s: $%s\n", itemName, price)
		} else {
			// 404
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "不存在商品%s\n", itemName)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不存在URL:%s\n", req.URL)
	}
}

func main() {
	db := database{
		"iphoneX":  9299,
		"iphone12": 10999,
		"nike":     299,
		"gucci":    599,
	}

	// database 实现了 ServeHTTP 接口，可以进行web响应处理
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
