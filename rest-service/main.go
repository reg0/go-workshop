package main

import (
	"github.com/labstack/echo/v4"
	"github.com/reg0/go-workshop/hello"
)

func main() {
	e := echo.New()
	e.GET("/hello", hello.SayHello)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/hello", hello.SayHello)
	//
	//err := http.ListenAndServe(":8080", mux)

	err := e.Start(":8080")

	if err != nil {
		panic(err)
	}
}