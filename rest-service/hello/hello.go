package hello

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelloMessage struct {
	Message string `json:"msg"`
}

func SayHello(c echo.Context) error {
	res := &HelloMessage{
		Message: "Hello",
	}

	return c.JSON(http.StatusOK, &res)
}