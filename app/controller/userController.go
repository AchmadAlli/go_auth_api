package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/middleware"
	"github.com/AchmadAlli/go_auth_api/app/request"
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
	g.GET("/me", c.user, middleware.AuthMiddleware())
	g.GET("/:id", c.show)
	g.POST("/", c.store)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.destroy)
}

func (c *UserController) index(ctx echo.Context) error {
	users, err := c.srv.Index()
	if err != nil {
		return helper.RestError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helper.RestApi(ctx, users)
}

func (c *UserController) store(ctx echo.Context) error {
	data, err := request.ValidateStoreUser(ctx)

	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := c.srv.Store(data)
	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}

	return helper.RestApi(ctx, user)
}

func (c *UserController) update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return helper.RestError(ctx, http.StatusNotFound, err.Error())
	}

	data, err := request.ValidateUpdateUser(ctx)

	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := c.srv.Update(data, uint(id))
	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}

	return helper.RestApi(ctx, user)
}

func (c *UserController) destroy(ctx echo.Context) error {
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

func (c *UserController) show(ctx echo.Context) error {
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

func (c *UserController) user(ctx echo.Context) error {
	id, isValid := ctx.Get("user_id").(uint)
	if !isValid {
		return helper.RestError(ctx, http.StatusUnauthorized, "Unauthorized")
	}

	user, err := c.srv.Show(uint(id))

	if err != nil {
		return helper.RestError(ctx, http.StatusInternalServerError, "")
	}

	return helper.RestApi(ctx, user)
}
