package request

import "github.com/labstack/echo"

type StoreUser struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
}

type Stores struct {
	UserName *string `json:"username"`
}

func ValidateStoreUser(ctx echo.Context) (*StoreUser, error) {
	req := new(StoreUser)
	err := ctx.Bind(req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *Stores) Val(c echo.Context) {

}
