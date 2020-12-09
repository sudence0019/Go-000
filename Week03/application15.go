package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 定义一个当前context
	pContext, cancel := context.WithCancel(context.Background())

	// 启动一个goroutine 传入当前context
	go func(pConcext context.Context) {
		// 根据传入的context生成一个子context
		cctx, _ := context.WithCancel(pContext)

		// 启动一个 gorotune
		go func(cctx context.Context) {
			for {
				select {
				case <-cctx.Done():
					fmt.Println("1-1-goroutine end")
					return
				default:
					fmt.Println("1-1-goroutine running")
					time.Sleep(5 * time.Second)
				}

			}
		}(cctx)
		go func(cctx context.Context) {
			for {
				select {
				case <-cctx.Done():
					fmt.Println("1-2-goroutine end")
					return
				default:
					fmt.Println("1-2-goroutine running")
					time.Sleep(5 * time.Second)
				}

			}
		}(cctx)
		for {
			select {
			case <-pContext.Done():
				fmt.Println("1-goroutine end")
				return
			default:
				fmt.Println("1-goroutine running")
				time.Sleep(5 * time.Second)
			}
		}
	}(pContext)

	time.Sleep(20 * time.Second)
	// 关闭parentContext，子context都会随之关闭。
	cancel()
	time.Sleep(4 * time.Second)
}
