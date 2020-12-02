package main

import (
	"fmt"
	"go-work/web"
)

func main() {
	users, code := web.GetAllUsers()
	if code == 200 {
		fmt.Println(len(users), cap(users))
		for i := 0; i < len(users); i++ {
			fmt.Println(users[i])
		}
	} else {
		fmt.Println("call api fail")
	}
}
