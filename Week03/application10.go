package main

import "fmt"

/**
 goroutine泄露
never start a goroutine without knowing when it will stop
*/
func leak() {
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("reveiver ", val)
	}()
}

func main() {

	leak()
}
