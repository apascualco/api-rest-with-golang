package api

import (
	"encoding/json"
	"net/http"
)

func HelloWorld() EndpointConfiguration {
	return EndpointConfiguration{
		PATH:     "/api/v1/hello",
		ROLES:    []string{"admin", "user"},
		METHODS:  []string{http.MethodGet},
		FUNCTION: hello,
	}
}

type Hi struct {
	Message string `json: "message"`
}

func hello(responseWriter http.ResponseWriter, request *http.Request) {
	hi := Hi{
		Message: "hello!!!!",
	}
	json.NewEncoder(responseWriter).Encode(hi)
}
