package main

import (
	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/database"
)

func main() {
	db, err := database.ConnectMysql()

	if err != nil {
		panic(err)
	}

	app := app.Init(db)
	app.Start()
}
