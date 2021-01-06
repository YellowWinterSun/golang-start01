package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//resp, err := http.Get("http://localhost:8081/hello")
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", b)
}
