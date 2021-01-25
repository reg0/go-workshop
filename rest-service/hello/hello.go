package hello

import (
	"github.com/labstack/echo/v4"
	"fmt"
	"log"
	"net/http"
)

type HelloMessage struct {
	Message string `json:"msg"`
}

func SayHello(c echo.Context) error {
	res := &HelloMessage{
		Message: "Hello",
	}

	log.Printf(fmt.Sprintf("Listening on port %v", 8080))

	return c.JSON(http.StatusOK, &res)
}