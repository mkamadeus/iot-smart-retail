package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mkamadeus/iot-smart-retail/app"
)

func main() {
	app, err := app.InitializeApp()
	if err != nil {
		panic(err)
	}

	app.Listen()
}
