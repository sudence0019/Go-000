package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {

		go func() {
			count := 1
			for true {
				count++
				if count%9999999 == 0 {
					fmt.Println("c1 runing....")
				}
			}
		}()

		time.Sleep(5 * time.Second)
		fmt.Println("B1 run overing")
	}()

	time.Sleep(10 * time.Second)
}
