package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	url := "http://httpbin:8080/headers"
	url += strings.Repeat("a", 1024*50)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
