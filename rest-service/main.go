package main

import (
	"net/http"
)

func sayHello(rw http.ResponseWriter, req *http.Request) {
	_, _ = rw.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(sayHello))

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}