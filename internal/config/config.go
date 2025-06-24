package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DbConNStr     string
}

func NewConfig() *Config {
	loadEnv()
	//log.Printf("sao khong duoc", os.Getenv("APP_PORT"))
	return &Config{

		ServerAddress: ":" + os.Getenv("APP_PORT"),
		DbConNStr:     os.Getenv("DB_CONN_STR"),
	}
}

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}
