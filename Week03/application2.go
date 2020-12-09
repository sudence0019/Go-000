package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "8080")
	})

	http.DefaultServeMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "8081")
	})

	// 使用协程启动一个服务 8081用于监控服务器
	go http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux)
	// 主协程启动8080处理应用请求
	http.ListenAndServe("127.0.0.1:8080", mux)
}
