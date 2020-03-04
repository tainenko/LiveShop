package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
	fmt.Println(user.Password)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}