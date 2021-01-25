package hello

import (
	"encoding/json"
	"net/http"
)

type HelloMessage struct {
	Message string `json:"msg"`
}

func SayHello(rw http.ResponseWriter, req *http.Request) {
	res := &HelloMessage{
		Message: "Hello",
	}

	bytes, _ := json.Marshal(res)

	rw.Header()["content-type"] = []string{"application/json"}
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(bytes)
}