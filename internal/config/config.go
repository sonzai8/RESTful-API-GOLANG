package config

import "os"

type Config struct {
	ServerAddress string
	DbConNStr     string
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: ":" + os.Getenv("8084"),
		DbConNStr:     os.Getenv("DB_CONN_STR"),
	}
}
