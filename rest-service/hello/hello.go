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

func logInfo(format string, args ...interface{}) {
	log.Printf(fmt.Sprintf(format, args...))
}

func SayHello(c echo.Context) error {
	res := &HelloMessage{
		Message: "Hello",
	}

	logInfo("Listening on port %v", 8080)
	return c.JSON(http.StatusOK, &res)
}