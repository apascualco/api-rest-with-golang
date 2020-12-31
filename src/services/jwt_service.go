package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func BuildToken(user string, roles []string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		TokenClaims{
			user,
			roles,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
				Issuer:    "apascualco",
			},
		},
	)
	return token.SignedString([]byte(getSecretOrDefualt()))
}

func getSecretOrDefualt() string {
	defaultValue := "secret"
	properties, err := ReadPropertiesFromPath("resources/application.properties")
	if err != nil {
		return defaultValue
	}
	secret := properties["token.secret"]
	if secret == "" {
		secret = defaultValue
	}
	return secret
}

type TokenBearer struct {
	Token string `json:"token"`
}

func ParseToken(userToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(getSecretOrDefualt()), nil
	})
	if token != nil {
		return token.Claims.(jwt.MapClaims), err
	}
	return nil, errors.New("Error psrsing token")
}
