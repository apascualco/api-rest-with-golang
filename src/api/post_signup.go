package api

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"apascualco.com/api-rest/src/sqlite3/repository"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup() EndpointConfiguration {
	return EndpointConfiguration{
		PATH:     "/api/v1/signup",
		METHODS:  []string{http.MethodPost},
		FUNCTION: signup,
	}
}

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func signup(responseWriter http.ResponseWriter, request *http.Request) {
	user := &User{}

	if json.NewDecoder(request.Body).Decode(user) != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := mail.ParseAddress(user.User)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	hashedPasswpord, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	userEntity := &repository.UserEntity{}
	userEntity.Password = string(hashedPasswpord)
	userEntity.Email = address.Address

	if userEntity.Create() != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
