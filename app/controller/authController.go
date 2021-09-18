package controller

import (
	"net/http"

	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/request"
	"github.com/AchmadAlli/go_auth_api/app/service"
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/labstack/echo"
)

type AuthController struct {
	srv *service.AuthService
}

func ListenAuth(app *app.App) {
	g := app.E.Group("/auth")
	c := AuthController{service.CreateAuthService(app)}

	g.POST("/login", c.login)
}

func (c *AuthController) login(ctx echo.Context) error {
	auth, err := request.ValidateAuth(ctx)

	if err != nil {
		return helper.RestError(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := c.srv.Login(auth)
	if err != nil {
		return helper.RestError(ctx, http.StatusUnauthorized, err.Error())
	}

	return helper.RestApi(ctx, user)
}
