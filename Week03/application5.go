package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// never start a goroutine without knowning when it will stop

func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func serverApp5() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "8080")
	})
	return http.ListenAndServe("localhost:8080", mux)
}

var count5 int

func serverDebug5() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "8081")
		fmt.Fprintf(writer, "8081")
		count5++
		if count5 > 10 {
			panic(errors.New("数量超了"))
		}

	})
	return http.ListenAndServe("localhost:8081", http.DefaultServeMux)
}

type H1 struct {
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})

	//go func() {
	//	done <-serverDebug5()
	//}()
	//
	//go func() {
	//	done <-serverApp5()
	//}()

	stoped := false
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error:", err)
		}
		if !stoped {
			stoped = true
			close(stop)
		}
	}
}
