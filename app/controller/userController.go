package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/service"
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type UserController struct {
	srv *service.UserService
}

func ListenUser(app *app.App) {
	g := app.E.Group("/api/users")
	c := UserController{service.CreateUserService(app)}

	g.GET("", c.index)
	g.GET("/", c.index)
	g.GET("/:id", c.show)
	g.POST("/", c.store)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.destroy)
}

func (c UserController) index(ctx echo.Context) error {
	users, err := c.srv.Index()
	if err != nil {
		return helper.RestError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helper.RestApi(ctx, users)
}

func (c UserController) store(ctx echo.Context) error {
	return helper.RestApi(ctx, "user stored!")
}

func (c UserController) update(ctx echo.Context) error {
	return helper.RestApi(ctx, "user updated!")
}

func (c UserController) destroy(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return helper.RestError(ctx, http.StatusNotFound, err.Error())
	}

	_, err = c.srv.Destroy(uint(id))
	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}

	return helper.RestApi(ctx, nil)
}

func (c UserController) show(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return helper.RestError(ctx, http.StatusNotFound, err.Error())
	}

	user, err := c.srv.Show(uint(id))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.RestError(ctx, http.StatusNotFound, err.Error())
	}

	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}
	return helper.RestApi(ctx, user)
}
