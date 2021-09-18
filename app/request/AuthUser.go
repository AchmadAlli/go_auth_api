package request

import "github.com/labstack/echo"

type AuthUser struct {
	Username string `json:"username" validate:"requried"`
	Password string `json:"password" validate:"requried"`
}

func ValidateAuth(ctx echo.Context) (*AuthUser, error) {
	req := new(AuthUser)
	err := ctx.Bind(req)

	if err != nil {
		return nil, err
	}

	return req, nil
}
