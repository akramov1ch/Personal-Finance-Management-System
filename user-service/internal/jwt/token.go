package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("dilshod333")

func CreateToken(email string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func VerifyToken(bearerToken string) (string, int64, error) {
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", 0, fmt.Errorf("token must start with 'Bearer '")
	}

	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", 0, err
	}


	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return "", 0, fmt.Errorf("token has expired")
			}
			email, ok := claims["email"].(string)
			if !ok {
				return "", 0, fmt.Errorf("invalid token: missing email claim")
			}
			return email, int64(exp), nil
		}
	}

	return "", 0, fmt.Errorf("invalid token")
}
