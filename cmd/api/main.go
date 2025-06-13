package main

import (
	"main/internal/app"
	"main/internal/config"
)

func main() {
	cfg := config.NewConfig()
	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
