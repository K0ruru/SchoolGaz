package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userNIS int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	
	claims := token.Claims.(jwt.MapClaims)
	claims["user_nis"] = userNIS
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

	tokenString, err := token.SignedString([]byte("heheboay"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}