package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("test")

type JWTClaims struct {
	EmployeeId uint   `json:"employee_id"`
	Role       string `json:"job_role"`
	jwt.StandardClaims
}

func GenerateJWT(employeeId uint, role string) (string, error) {
	claims := JWTClaims{
		EmployeeId: employeeId,
		Role:       role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ParseJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	return claims, nil
}
