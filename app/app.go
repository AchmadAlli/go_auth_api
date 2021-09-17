package app

import (
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type App struct {
	E  *echo.Echo
	DB *gorm.DB
}

func Init(db *gorm.DB) App {
	app := App{
		E:  echo.New(),
		DB: db,
	}

	app.E.GET("/", healthCheck)

	return app
}

func (app *App) Start() {
	app.E.Start(":8000")
}

func healthCheck(ctx echo.Context) error {
	return helper.RestApi(ctx, "")
}
