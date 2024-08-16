package Infrastructure

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"

)

var jwtSecret = []byte("jwt_secret_key0")

func GenerateToken(userID string, userName string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userID,
		"userName": userName,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// Handle error or invalid token
	if err != nil || !token.Valid {
		return nil, err
	}

	// Return the claims as a MapClaims
	return token.Claims.(jwt.MapClaims), nil
}
