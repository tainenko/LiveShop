package databases

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUserTable() {
	DB.Query(`
		CREATE TABLE IF NOT EXISTS users( 
			id serial PRIMARY KEY,
			name VARCHAR (100) NOT NULL,
			password VARCHAR (355) NOT NULL,
			email VARCHAR (355) UNIQUE NOT NULL,
			created_on TIMESTAMP NOT NULL default current_timestamp,
			updated_at TIMESTAMP NOT NULL default current_timestamp 
			)`,
	)
}

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

