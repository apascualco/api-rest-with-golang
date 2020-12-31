package server

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"apascualco.com/api-rest/src/services"
)

var PROPERTY_PREFIX = "rest.endpoint.autenticatio"

type Middleware struct {
	ROLES map[string][]string
}

func NewMiddleware(rolesByPath map[string][]string) Middleware {
	return Middleware{
		ROLES: rolesByPath,
	}
}

func (middleware *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {

		endpointRoles := middleware.ROLES[request.URL.Path]

		if len(endpointRoles) > 0 {
			userRoles, err := getUserRoleByRequest(request)
			if err != nil && !anyRoleMatch(endpointRoles, userRoles) {
				responseWriter.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(responseWriter, request)

	})
}

func getUserRoleByRequest(request *http.Request) ([]interface{}, error) {

	userToken := getTokenFromRequest(request)
	claims, err := services.ParseToken(userToken)
	if err != nil {
		return nil, err
	}
	expireTokenTime := parseFloat64ToTime(claims["exp"].(float64))
	userRoles := claims["roles"].([]interface{})
	if expireTokenTime.Before(time.Now()) || len(userRoles) == 0 {
		return nil, errors.New("Error parsing token")
	}
	return userRoles, nil
}

func anyRoleMatch(endpointRoles []string, userRoles []interface{}) bool {
	for _, endpointRole := range endpointRoles {
		for _, userRole := range userRoles {
			if strings.EqualFold(endpointRole, fmt.Sprintf("%v", userRole)) {
				return true
			}
		}
	}
	return false
}

func getTokenFromRequest(request *http.Request) string {
	authorization := request.Header.Get("Authorization")
	if authorization != "" {
		authorizationParts := strings.Fields(authorization)
		if len(authorizationParts) == 2 {
			tokenType := authorizationParts[0]
			token := authorizationParts[1]
			if strings.EqualFold(tokenType, "bearer") && len(token) > 0 {
				return token
			}
		}
	}
	return ""
}

func parseFloat64ToTime(expreTime float64) time.Time {
	sec, dec := math.Modf(expreTime)
	return time.Unix(int64(sec), int64(dec*(1e9)))
}
