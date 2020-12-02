package web

import (
	"fmt"
	"go-work/entity"
	"go-work/service"
)

func GetAllUsers() ([]*entity.User, int) {
	users, err := service.GetAllUserService()
	if err != nil {
		fmt.Printf("err: %+v", err)
		return nil, 500
	}
	return users, 200
}
