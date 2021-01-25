package main

import (
	"net/http"
)

func sayHello(rw http.ResponseWriter, req *http.Request) {
	_, _ = rw.Write([]byte("hello"))
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(sayHello))

	if err != nil {
		panic(err)
	}
}