package controller

import (
	"net/http"

	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/service"
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/labstack/echo"
)

type UserSrv struct {
	srv *service.UserService
}

func ListenUser(app *app.App) {
	g := app.E.Group("/api/users")
	c := UserSrv{service.CreateUserService(app)}

	g.GET("", c.index)
	g.GET("/", c.index)
	g.GET("/:id", c.show)
	g.POST("/", c.store)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.destroy)
}

func (s UserSrv) index(ctx echo.Context) error {
	users, err := s.srv.Index()
	if err != nil {
		return helper.RestError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helper.RestApi(ctx, users)
}

func (s UserSrv) store(ctx echo.Context) error {
	return helper.RestApi(ctx, "user stored!")
}

func (s UserSrv) update(ctx echo.Context) error {
	return helper.RestApi(ctx, "user updated!")
}

func (s UserSrv) destroy(ctx echo.Context) error {
	return helper.RestApi(ctx, "user destroyed!")
}

func (s UserSrv) show(ctx echo.Context) error {
	return helper.RestApi(ctx, "user")
}
