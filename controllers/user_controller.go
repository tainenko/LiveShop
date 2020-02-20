package controllers

import (
	"fmt"
	"github.com/LiveShop/databases"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Register(c *gin.Context){
	var user databases.RegisterInfo
	c.Bind(&user)
	fmt.Println(user)
	validErr:=ValidateUser(user,[]string{})
	exist := isUserExist(user)
	if exist == true{
		validErr=append(validErr,"Email already exist!")
	}
	fmt.Println(validErr)
	if len(validErr)>0{
		c.JSON(http.StatusUnprocessableEntity,gin.H{"success":false,"errors":validErr})
		return
	}
	databases.HashPassword(&user)
	_,err:=databases.DB.Query(databases.CreateUserQuery,user.Name,user.Password,user.Email)
	if err!=nil{
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created succesfully"})
}

func isUserExist(user databases.RegisterInfo) bool{
	rows,err:=databases.DB.Query(databases.IsUserExist,user.Email)
	if err!=nil{
		return false
	}
	if !rows.Next(){
		return false
	}
	return true
}

func ValidateUser(user databases.RegisterInfo,err []string) []string{
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	emailCheck:=regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck!=true{
		err=append(err,"Invalid email")
	}
	if len(user.Password)<=4{
		err=append(err,"Invalid Password. Password should have more than 4 charaters.")
	}
	if len(user.Name)<1{
		err=append(err,"Invalid name.")
	}
	fmt.Printf("Pass")
	return err
}