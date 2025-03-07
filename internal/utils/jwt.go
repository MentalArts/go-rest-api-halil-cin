package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("TESTB4ENV")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokens(username string, role string) (string, string, error) {
	// Access token
	accessTokenExp := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExp.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Refresh token
	refreshTokenExp := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExp.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
