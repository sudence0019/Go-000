package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go-work/entity"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/mybatis?charset=utf8")
	if err != nil {
		panic(err)
	}
	Db = db
}
func QueryAllUser() ([]*entity.User, error) {
	rows, err := Db.Query("select id,name as name,class_id as classId from student")
	if err != nil {
		return nil, errors.Wrap(err, "query fail")
	}
	defer rows.Close()
	users := make([]*entity.User, 10, 10)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.Id, &user.Name, &user.ClassId)
		if err != nil {
			return nil, errors.Wrap(err, "format user fail")
		}
		users = append(users, user)
	}
	return users, nil

}
