package main

import (
	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/controller"
	"github.com/AchmadAlli/go_auth_api/database"
)

func main() {
	db, err := database.ConnectMysql()

	if err != nil {
		panic(err)
	}

	app := app.Init(db)
	listenServices(&app)
	app.Start()
}

func listenServices(app *app.App) {
	controller.ListenUser(app)
}
