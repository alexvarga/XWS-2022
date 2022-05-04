package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secret = "this is a secret"

type TokenClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateJWT(userId string, username string) (string, error) {
	claims := TokenClaims{UserId: userId, Username: username, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err, "jwt error")
		return "", err
	}
	return tokenString, nil

}
