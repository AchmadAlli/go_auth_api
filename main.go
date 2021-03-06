package main

import (
	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/controller"
	"github.com/AchmadAlli/go_auth_api/app/model"
	"github.com/AchmadAlli/go_auth_api/database"
	"github.com/AchmadAlli/go_auth_api/database/seeder"
	"github.com/jinzhu/gorm"
)

func main() {
	database.CreateDatabase("try_golang")
	db, err := database.ConnectMysql()

	if err != nil {
		panic(err)
	}

	migrate(db)
	seed(db)

	app := app.Init(db)
	listenServices(&app)
	app.Start()
}

func listenServices(app *app.App) {
	controller.ListenUser(app)
	controller.ListenAuth(app)
}

func migrate(db *gorm.DB) {
	database.HandleMigration(db, &model.User{})
}

func seed(db *gorm.DB) {
	seeder.SeedUser(db)
}
