package config

import (
	"github.com/go-chi/jwtauth/v5"
	"os"
)

var TokenAuth *jwtauth.JWTAuth

func InitJwt() {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	TokenAuth = jwtauth.New("HS256", []byte(jwtSecretKey), nil)
}
