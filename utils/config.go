package utils

import (
	"github.com/joho/godotenv"
	"log"
)


func LoadConfig() map[string]string {
	var myEnv map[string]string
	var err error
	myEnv, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return myEnv
}
