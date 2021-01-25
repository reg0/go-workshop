package main

import (
	"github.com/reg0/go-workshop/hello"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello.SayHello)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}