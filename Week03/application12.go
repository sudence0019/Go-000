package main

import (
	"fmt"
	"time"
)

var countApplication12 = 0
var flag = false

func main() {
	go func() {

		for !flag {
			time.Sleep(2 * time.Second)
			countApplication12++
			fmt.Println(countApplication12)

		}
		fmt.Println("g1 over")

	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			if countApplication12 == 20 {
				fmt.Println("end .....")
				flag = true
				break
			}
			fmt.Println("g2......", countApplication12)

		}
	}()

	time.Sleep(20000 * time.Second)
}
