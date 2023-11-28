package utils

import (
	"github.com/hertz-contrib/jwt"
)

var identityKey = "FanR"

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	jwt.HertzJWTMiddleware
}

func GenerateToken(id uint, userName string) (string, error) {
	return "nil", nil
}
