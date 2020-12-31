package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Handler() http.Handler
}

func BuildServerRoutes() Server {
	api := &api{}

	router := mux.NewRouter()

	rolesByPath := make(map[string][]string)
	applicationEndpoints := NewApplicationEndpoints()
	for _, endpointConfiguration := range applicationEndpoints.ENDPOINTS {
		router.HandleFunc(endpointConfiguration.PATH, endpointConfiguration.FUNCTION).Methods(endpointConfiguration.METHODS...)
		rolesByPath[endpointConfiguration.PATH] = endpointConfiguration.ROLES
	}

	middleware := NewMiddleware(rolesByPath)
	router.Use(middleware.Middleware)
	api.router = router
	return api
}

func (api *api) Handler() http.Handler {
	return api.router
}
