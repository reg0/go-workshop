package main

import (
	"github.com/labstack/echo/v4"
	"github.com/reg0/go-workshop/hello"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/hello", echo.WrapHandler(http.HandlerFunc(hello.SayHello)))

	//mux := http.NewServeMux()
	//mux.HandleFunc("/hello", hello.SayHello)
	//
	//err := http.ListenAndServe(":8080", mux)

	err := e.Start(":8080")

	if err != nil {
		panic(err)
	}
}