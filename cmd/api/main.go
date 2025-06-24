package main

import (
	"log"
	"main/internal/app"
	"main/internal/config"
)

func main() {
	cfg := config.NewConfig()
	log.Printf("Starting server... %s", cfg.ServerAddress)
	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
