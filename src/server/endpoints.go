package server

import (
	endpoint_api "apascualco.com/api-rest/src/api"
)

type ApplicationEndpoints struct {
	ENDPOINTS []endpoint_api.EndpointConfiguration
}

func (applicationEndpoints *ApplicationEndpoints) add(endpointConfiguration endpoint_api.EndpointConfiguration) {
	applicationEndpoints.ENDPOINTS = append(applicationEndpoints.ENDPOINTS, endpointConfiguration)
}

func NewApplicationEndpoints() ApplicationEndpoints {
	applicationEndpoints := &ApplicationEndpoints{ENDPOINTS: []endpoint_api.EndpointConfiguration{}}

	applicationEndpoints.add(endpoint_api.PostLoginJwt())
	applicationEndpoints.add(endpoint_api.UserSignup())
	applicationEndpoints.add(endpoint_api.HelloWorld())

	return *applicationEndpoints
}
