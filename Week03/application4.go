package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverApp2() {
	defer func() {
		fmt.Println("serverApp2 退出了")
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "app")
	})
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {

		log.Fatal(err)
	}
}

var count2 int

func serverDebug2() {
	defer func() {
		fmt.Println("serverDebug2 退出了")
	}()
	server := http.NewServeMux()
	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "debuger")
		if count2 > 10 {
			log.Fatal(count2)

		}
		count2++
	})
	if err := http.ListenAndServe("localhost:8081", server); err != nil {
		fmt.Println("出现异常了")
		log.Fatal(err)
		// 使用log.Fatal调用了os.Exit 无条件终止程序，defer不会调用。
	}
}

func main() {
	go serverApp2()
	go serverDebug2()
	select {}
}
