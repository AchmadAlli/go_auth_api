package helper

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type apiContract struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Missing env data")
	}

	return os.Getenv(key)
}

func RestApi(ctx echo.Context, data interface{}) error {

	res := apiContract{
		Status: "ok",
		Data:   data,
		Error:  nil,
	}

	return ctx.JSON(http.StatusOK, res)
}

func RestError(ctx echo.Context, code int, msg string) error {

	res := apiContract{
		Status: "error",
		Data:   nil,
		Error:  msg,
	}

	return ctx.JSON(code, res)
}
