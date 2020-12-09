package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tr := NewTracker()
	go tr.Run()
	// 发送事件
	tr.Event(context.Background(), "test1")
	tr.Event(context.Background(), "test1")
	tr.Event(context.Background(), "test1")

	// 创建一个context
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(8*time.Second))
	// 关闭ctx
	defer cancel()
	// 关闭tracker
	tr.Shutdown(ctx)

}

// tracker结构体
type Tracker struct {
	//数据channel
	ch chan string
	// 结束标志
	stop chan struct{}
}

// 构造函数
func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

// 发送消息
func (t *Tracker) Event(ctx context.Context, data string) error {
	//使用select 不会阻塞，同时调用者不用开启一个goroutine
	select {
	// 有数据来了放到ch里就返回
	case t.ch <- data:
		return nil
	// 如果超时，ctx结束就返回error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	// 循环取数据
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	// ch关闭
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	// 关闭channel
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}
