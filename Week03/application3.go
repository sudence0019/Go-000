package main

import (
	"errors"
	"fmt"
	"net/http"
)

func serverApp() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "8080")
	})
	return http.ListenAndServe("localhost:8080", mux)
}

var count int

func serverDebug() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "8081")
		fmt.Fprintf(writer, "8081")
		count++
		if count > 10 {
			panic(errors.New("数量超了"))
		}

	})
	return http.ListenAndServe("localhost:8081", http.DefaultServeMux)
}
func main() {
	// 由调用者决定是否有启用协程运行函数。而不是函数内部启动一个野生的线程。
	// 如果go serverDebug()挂掉了，serverApp()还在继续运行。
	go serverDebug()
	serverApp()
}
