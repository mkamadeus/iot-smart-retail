package main

import (
	"github.com/joho/godotenv"
	"github.com/mkamadeus/iot-smart-retail/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app, err := app.InitializeApp()
	if err != nil {
		panic(err)
	}

	app.Listen()
}
