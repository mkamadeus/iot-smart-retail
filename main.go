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

	a, err := app.GetApp()
	if err != nil {
		panic(err)
	}
	a.Listen(":8080")
}
