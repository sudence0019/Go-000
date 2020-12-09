package main

import (
	"context"
	"fmt"
)

func main() {
	p1 := context.Background()
	fmt.Println(p1.Value("1"))
	p2 := context.WithValue(p1, "1", "1")
	fmt.Println(p2 == p1)
	fmt.Println(p2)
	fmt.Println(p2.Value("1"))

	p3, _ := context.WithCancel(p2)
	fmt.Println(p3)
	fmt.Println(p3.Value("1"))
}
