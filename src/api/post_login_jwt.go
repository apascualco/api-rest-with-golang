package api

import (
	"encoding/json"
	"log"
	"net/http"

	"apascualco.com/api-rest/src/services"
	"apascualco.com/api-rest/src/sqlite3/repository"
	"golang.org/x/crypto/bcrypt"
)

func PostLoginJwt() EndpointConfiguration {
	return EndpointConfiguration{
		PATH:     "/api/v1/login",
		METHODS:  []string{http.MethodPost},
		FUNCTION: login,
	}
}

type Credentials struct {
	Password string `json:"password"`
	User     string `json:"user"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func login(responseWriter http.ResponseWriter, request *http.Request) {
	credentials := &Credentials{}

	if json.NewDecoder(request.Body).Decode(credentials) != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	userEntity, err := validateCredentials(credentials)
	if err != nil {
		responseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenResponse := &TokenResponse{}

	token, err := services.BuildToken(userEntity.Email, []string{"admin", "luser"})
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenResponse.Token = token
	json.NewEncoder(responseWriter).Encode(tokenResponse)
}

func validateCredentials(credentials *Credentials) (repository.UserEntity, error) {
	userEntity := &repository.UserEntity{}
	err := userEntity.GetByEmail(credentials.User)
	if err != nil {
		return *userEntity, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(credentials.Password))
	return *userEntity, err
}
