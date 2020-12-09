package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "监控1")
	go watch(ctx, "监控2")
	go watch(ctx, "监控3")
	time.Sleep(40 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)

	//方法是获取设置的截止时间，时间，是否到期。
	// ctx.Deadline()

	// done方法返回一个只读的chan,类型是struct{},我们在goroutine中，如果该方法返回的chan可以读取，意味着parentContext已经发起了取消请求，我们收到这个信号，做清理操作。
	// ctx.Done()

	// err方法返回取消的错误原因，因为什么context取消。
	// ctx.Err()

	// value获取context上返回的值
	// ctx.Value("key")
}

func watch(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控退出了")
			return
		default:
			fmt.Println(name, "监控运行中")
			time.Sleep(5 * time.Second)
		}
	}
}
