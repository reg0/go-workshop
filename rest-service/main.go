package main

import (
	"encoding/json"
	"net/http"
)

type HelloMessage struct {

}

func sayHello(rw http.ResponseWriter, req *http.Request) {
	res := &HelloMessage{}

	bytes, _ := json.Marshal(res)

	_, _ = rw.Write(bytes)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}