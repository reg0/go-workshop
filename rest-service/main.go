package main

import (
	"net/http"
)

type Handler struct {
}

func (s *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_, _ = rw.Write([]byte("hello"))
}

func main() {
	err := http.ListenAndServe(":8080", &Handler{})

	if err != nil {
		panic(err)
	}
}