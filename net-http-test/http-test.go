package net_http_test

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func NetRun() {
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":8080", nil)
	// http请求
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取返回的数据
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印返回的数据
	fmt.Println(len(content))
}
