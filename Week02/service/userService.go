package service

import (
	userDao "go-work/dao"
	"go-work/entity"
)

func GetAllUserService() ([]*entity.User, error) {
	return userDao.QueryAllUser()
}
