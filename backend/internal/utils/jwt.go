package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
		ErrInvalidSigningMethod		= errors.New("Invalid signing method")
		ErrInvalidToken				= errors.New("Invalid token")
)

type Claims struct {
	UserID   uint `json:"user_id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(GetJwtSecret())

func GetJwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret)
}

func GenerateToken(userID uint, email, fullname, role string) (string, error) {
	expiredTime := time.Now().Add(12 * time.Hour)
	
	claims := &Claims{
		UserID:   userID,
		Email:    email,
		Fullname: fullname,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidationToken(tokenString string) (*Claims, error){
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}