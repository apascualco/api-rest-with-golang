package api

import "net/http"

type EndpointConfiguration struct {
	PATH     string
	METHODS  []string
	ROLES    []string
	FUNCTION func(responseWriter http.ResponseWriter, request *http.Request)
}
