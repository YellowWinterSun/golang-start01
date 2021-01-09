package main

import (
	"fmt"
	"log"
	"net/http"
)

var requestCount int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = path[1:]
	fmt.Println("请求的命令:", path)

	if path == "hello" {
		requestCount++
	} else if path == "count" {
		fmt.Fprintf(w, "count = %d\n", requestCount)
		return
	}

	fmt.Fprintf(w, "URL.Path = %q\n", path)
}
