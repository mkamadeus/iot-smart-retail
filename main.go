package main

import (
	"github.com/joho/godotenv"
	"github.com/mkamadeus/iot-smart-retail/app"
)

func main() {
	// dotenv
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	a, err := app.InitializeApp()
	if err != nil {
		panic(err)
	}
	a.Server.Listen(":8080")
}
