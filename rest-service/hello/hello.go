package hello

import (
	"encoding/json"
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

	bytes, _ := json.Marshal(res)

	rw := c.Response()
	rw.Header()["content-type"] = []string{"application/json"}
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(bytes)

	return nil
}