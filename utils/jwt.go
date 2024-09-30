package utils

import "github.com/golang-jwt/jwt/v5"

var jwtKey = "SECRET"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return webToken, nil
}
