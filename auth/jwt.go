package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const hello string = "key"

type JwtCustomClaim struct {
	ID uuid.UUID
	jwt.StandardClaims
}

func GenAccessToken(ctx context.Context, userID uuid.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString([]byte(hello))
	if err != nil {
		return "", err
	}

	return token, nil
}
func GenRfreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString([]byte(hello))
	if err != nil {
		return "", err
	}

	return token, nil
}
func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return []byte(hello), nil
	})
}
