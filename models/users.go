package models

import (
	"fmt"
	"github.com/LiveShop/databases"
	"github.com/LiveShop/utils"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	CreateOn string `json:"created_on"`
	UpdateAt string `json:"update_at"`
}

type RegisterInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HashPassword(user *RegisterInfo) {
	user.Password = string(utils.MD5(user.Password))
	fmt.Println(user.Password)
}

func QueryUserId(email, password string) int {
	row := databases.DB.QueryRow(databases.QueryUserWithParam, email, password)
	id := 0
	row.Scan(&id)
	return id
}
