package request

import "github.com/labstack/echo"

type UpdateUser struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func ValidateUpdateUser(ctx echo.Context) (*UpdateUser, error) {
	req := new(UpdateUser)
	err := ctx.Bind(req)

	if err != nil {
		return nil, err
	}

	return req, nil
}
