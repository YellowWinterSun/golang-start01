package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func (db database) getDollars(w http.ResponseWriter, req *http.Request) {
	itemName := req.URL.Query().Get("item")
	if price, ok := db[itemName]; ok {
		fmt.Fprintf(w, "存在商品%s: %s\n", itemName, price)
	} else {
		// 404
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不存在商品%s\n", itemName)
	}
}

func main() {
	db := database{
		"iphoneX":  9299,
		"iphone12": 10999,
		"nike":     299,
		"gucci":    599,
	}

	/*
		 如果我们的应用只需要一个Web服务器
		 可以利用http包下的DefaultServeMux，这个是全局的。一个进程对应一个，相当于定义全局的mux。
		而Web服务器3的例子中，如
		<p>
			//mux := http.NewServeMux()
			//mux.Handle("/list", http.HandlerFunc(db.list))
			//mux.Handle("/getDollars", http.HandlerFunc(db.getDollars))
		</p>
		这个是定义一个mux，可以定义多个mux，但是弊端就是，我们加 path映射方法时，不方便
	*/
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/getDollars", db.getDollars)

	// nil相当于用默认的全局DefaultServeMux
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
