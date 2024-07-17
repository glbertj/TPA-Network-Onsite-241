package utils

import (
	"back-end/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secret_key"

func GenerateJWT(user model.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": user.UserId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetJWTClaims(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return "", ParseTokenError
	}

	if !token.Valid {
		fmt.Println("Token is not valid")
		return "", TokenInvalid
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		fmt.Println("Error asserting claims")
		return "", ClaimInvalid
	}

	issuer, ok := (*claims)["iss"].(string)
	if !ok {
		fmt.Println("Error extracting issuer claim")
		return "", ClaimInvalid
	}

	return issuer, nil

}
