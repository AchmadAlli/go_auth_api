package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type AuthToken struct {
	UserID uint
}

func AuthMiddleware() echo.MiddlewareFunc {
	return authHandler
}

func authHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, err := validate(ctx)
		if err != nil {
			return ctx.NoContent(http.StatusUnauthorized)
		}

		ctx.Set("user_id", token.UserID)

		return next(ctx)
	}
}

func validate(ctx echo.Context) (*AuthToken, error) {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "missing token")
	}

	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})

	if err != nil {
		return nil, err
	}

	if !payload.Valid {
		return nil, err
	}

	claims := payload.Claims.(jwt.MapClaims)
	id, isValid := claims["user_id"].(float64)

	if !isValid {
		return nil, err
	}

	return &AuthToken{uint(id)}, nil
}
