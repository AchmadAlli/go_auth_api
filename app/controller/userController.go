package controller

import (
	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/labstack/echo"
)

func ListenUser(app *app.App) {
	g := app.E.Group("/api/users")

	g.GET("", index)
	g.GET("/", index)
	g.GET("/:id", show)
	g.POST("/", store)
	g.PUT("/:id", update)
	g.DELETE("/:id", destroy)
}

func index(ctx echo.Context) error {
	return helper.RestApi(ctx, "users")
}

func store(ctx echo.Context) error {
	return helper.RestApi(ctx, "user stored!")
}

func update(ctx echo.Context) error {
	return helper.RestApi(ctx, "user updated!")
}

func destroy(ctx echo.Context) error {
	return helper.RestApi(ctx, "user destroyed!")
}

func show(ctx echo.Context) error {
	return helper.RestApi(ctx, "user")
}
