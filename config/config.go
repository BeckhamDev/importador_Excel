package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ConnectDBUrl = ""
	SecretKey = ""
)

func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ConnectDBUrl = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PASSWORD"))

	SecretKey = os.Getenv("SECRET_KEY")
}
