package app

import (
    "sync"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type App struct {
    AppServer    *fiber.App
    DBConnection *gorm.DB
    // MQTTClient   *mqtt.Client
}

var retailApp *App
var applock = &sync.Mutex{}

func GetApp() (*App, error) {
    var err error
    if retailApp == nil {
        applock.Lock()
        defer applock.Unlock()
        app := InitializeFiberApp()

        db, err := InitializeDB()
        if err != nil {
            return nil, err
        }

        retailApp = &App{
            AppServer:    app,
            DBConnection: db,
        }
    }
    return retailApp, err
}

func (app App) Listen(port string) {
    app.AppServer.Listen(port)
}
