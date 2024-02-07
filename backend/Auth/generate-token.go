package auth

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

func CreateToken(nis int) (string, error) {
    // Create a new JWT token with HMAC signing method and secret key
    token := jwt.New(jwt.SigningMethodHS256)

    // Define the claims to be added to the token (in this case, nis)
    claims := token.Claims.(jwt.MapClaims)
    claims["nis"] = nis
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set token expiration time (e.g., 1 day)

    // Create a signed token using your secret key
    tokenString, err := token.SignedString([]byte("RahasiaBangGabolehLiat"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

